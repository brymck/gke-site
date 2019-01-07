package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (fe *frontendServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	vars := mux.Vars(r)
	name := vars["name"]

	message, err := fe.sayHello(r.Context(), name)
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to say hello to " + name), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Error(err)
	}
}

func (fe *frontendServer) squareHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	vars := mux.Vars(r)
	numberText := vars["number"]
	number64, err := strconv.ParseFloat(numberText, 32)
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to convert " +numberText+ " to a number"), http.StatusInternalServerError)
		return
	}
	number32 := float32(number64)

	message, err := fe.getSquare(r.Context(), number32)
	if err != nil {
		message := fmt.Sprintf("failed to square %f", number32)
		renderHTTPError(log, r, w, errors.Wrap(err, message), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Error(err)
	}
}


func renderHTTPError(log logrus.FieldLogger, _ *http.Request, w http.ResponseWriter, err error, code int) {
	log.WithField("error", err).Error("request error")
	w.WriteHeader(code)
}
