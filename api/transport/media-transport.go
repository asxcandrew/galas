package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/asxcandrew/galas/api/endpoint"
	"github.com/asxcandrew/galas/social/media"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeMediaHandler(s media.MediaService, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getMediaHandler := kithttp.NewServer(
		endpoint.MakeGetMediaEndpoint(s),
		decodeGetMediaRequest,
		encodeResponse,
		opts...,
	)

	createMediaHandler := kithttp.NewServer(
		endpoint.MakeCreateMediaEndpoint(s),
		decodeCreateMediaRequest,
		encodeResponse,
		opts...,
	)

	deleteMediaHandler := kithttp.NewServer(
		endpoint.MakeGetMediaEndpoint(s),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/api/v1/media/{uuid}", getMediaHandler).Methods("GET")
	r.Handle("/api/v1/media", createMediaHandler).Methods("POST")
	r.Handle("/api/v1/media/{uuid}", deleteMediaHandler).Methods("DELETE")

	return r
}

func decodeGetMediaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	val, ok := vars["username"]

	if !ok {
		return nil, errBadRoute
	}

	return endpoint.GetUserRequest{Username: val}, nil
}

func decodeCreateMediaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = endpoint.CreateMediaRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}
