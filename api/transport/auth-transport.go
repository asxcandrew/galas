package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/asxcandrew/galas/workers"

	"github.com/asxcandrew/galas/api/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeAuthHandler(w workers.AuthWorker, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	loginHandler := kithttp.NewServer(
		endpoint.MakeLoginEndpoint(w),
		decodeLoginRequest,
		encodeResponse,
		opts...,
	)

	registerHandler := kithttp.NewServer(
		endpoint.MakeRegisterEndpoint(w),
		decodeRegisterRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/api/v1/auth/login", loginHandler).Methods("POST")
	r.Handle("/api/v1/auth/register", registerHandler).Methods("POST")

	return r
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = endpoint.LoginRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = endpoint.RegisterRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}
