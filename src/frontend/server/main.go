package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"google.golang.org/grpc"
)

const (
	port            = "8080"
	cookieMaxAge    = 60 * 60 * 48
	cookiePrefix    = "opto_"
	cookieSessionID = cookiePrefix + "session-id"
)

var (
	log *logrus.Logger
)

type ctxKeySessionID struct{}

type frontendServer struct {
	helloSvcAddr string
	helloSvcConn *grpc.ClientConn
	squareSvcAddr string
	squareSvcConn *grpc.ClientConn
}

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout
}

func main() {
	ctx := context.Background()
	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}
	addr := os.Getenv("LISTEN_ADDR")
	svc := new(frontendServer)
	mustMapEnv(&svc.helloSvcAddr, "HELLO_SERVICE_ADDR")
	mustMapEnv(&svc.squareSvcAddr, "SQUARE_SERVICE_ADDR")

	mustConnGRPC(ctx, &svc.helloSvcConn, svc.helloSvcAddr)
	mustConnGRPC(ctx, &svc.squareSvcConn, svc.squareSvcAddr)

	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", svc.helloHandler).Methods("GET")
	r.HandleFunc("/square/{number}", svc.squareHandler).Methods("GET")
	r.HandleFunc("/robots.txt", func(w http.ResponseWriter, _ *http.Request) {
		_, err := fmt.Fprint(w, "User-agent: *\nDisallow: /")
		if err != nil {
			log.Error(err)
		}
	})
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui/")))

	var handler http.Handler = r
	handler = &logHandler{log: log, next: handler}
	handler = ensureSessionID(handler)
	handler = &ochttp.Handler{Handler: handler, Propagation: &b3.HTTPFormat{}}

	log.Infof("starting server on " + addr + ":" + srvPort)
	log.Fatal(http.ListenAndServe(addr+":"+srvPort, handler))
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
