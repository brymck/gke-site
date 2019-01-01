package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (fe *frontendServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Debug("asking hello")

	message, err := fe.sayHello(r.Context(), "World")
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to say hello"), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Error(err)
	}
}

func renderHTTPError(log logrus.FieldLogger, r *http.Request, w http.ResponseWriter, err error, code int) {
	log.WithField("error", err).Error("request error")
	// errMsg := fmt.Sprintf("%+v", err)

	w.WriteHeader(code)
}
