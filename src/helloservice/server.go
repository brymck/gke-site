package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/brymck/gke-site/src/helloservice/genproto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	log *logrus.Logger

	port = flag.Int("port", 3550, "port to listen at")
)

func init() {
	log = logrus.New()
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

type greeter struct{}

func (p *greeter) GetGreeting(ctx context.Context, req *pb.GreetingRequest) (*pb.Greeting, error) {
	name := req.Name
	return &pb.Greeting{Message: "Hello, " + strings.Title(name) + "!"}, nil
}

func main() {
	flag.Parse()
	log.Infof("starting grpc server at :%d", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	svc := &greeter{}
	pb.RegisterHelloServiceServer(srv, svc)
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

