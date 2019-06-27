package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/asxcandrew/galas/api/endpoint"
	"github.com/asxcandrew/galas/social/user"
	"github.com/asxcandrew/galas/worker"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeUserHandler(s user.UserService, w worker.AuthWorker, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getUserHandler := kithttp.NewServer(
		endpoint.MakeGetUserEndpoint(s),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	)

	updateUserHandler := kithttp.NewServer(
		w.NewJWTParser(endpoint.MakeUpdateUserEndpoint(s)),
		decodePutUserRequest,
		encodeResponse,
		append(opts, kithttp.ServerBefore(worker.HTTPToContext()))...,
	)

	r := mux.NewRouter()

	r.Handle("/api/v1/users/{username:[a-z]}", getUserHandler).Methods("GET")
	r.Handle("/api/v1/users/{username:[a-z]}", updateUserHandler).Methods("PUT")

	return r
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	val, ok := vars["username"]

	if !ok {
		return nil, errBadRoute
	}

	return endpoint.GetUserRequest{Username: val}, nil
}

func decodePutUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	u, ok := vars["username"]

	if !ok {
		return nil, errBadRoute
	}

	req := endpoint.UpdateUserRequest{}

	req.Username = u

	if err := json.NewDecoder(r.Body).Decode(&req.Data); err != nil {
		return nil, err
	}

	return req, nil
}
