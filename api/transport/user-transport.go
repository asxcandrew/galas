package transport

import (
	"context"
	"net/http"

	"github.com/asxcandrew/galas/api/endpoint"
	"github.com/asxcandrew/galas/user"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeUserHandler(s user.UserService, logger log.Logger) http.Handler {
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

	r := mux.NewRouter()

	r.Handle("/api/v1/users/{username}", getUserHandler).Methods("GET")

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
