package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/asxcandrew/galas/api/representation"

	"github.com/asxcandrew/galas/api/endpoint"
	"github.com/asxcandrew/galas/item"
	"github.com/asxcandrew/galas/workers"
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var errBadRoute = errors.New("bad route")

func MakeItemHandler(s item.ItemService, w workers.AuthWorker, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	showItemHandler := kithttp.NewServer(
		endpoint.MakeShowItemEndpoint(s),
		decodeShowItemRequest,
		encodeResponse,
		opts...,
	)

	createItemHandler := kithttp.NewServer(
		w.NewJWTParser(endpoint.MakeCreateItemEndpoint(s)),
		decodeCreateItemRequest,
		encodeResponse,
		append(opts, kithttp.ServerBefore(gokitjwt.HTTPToContext()))...,
	)

	r := mux.NewRouter()

	r.Handle("/api/v1/item/{id}", showItemHandler).Methods("GET")
	r.Handle("/api/v1/item/", createItemHandler).Methods("POST")

	return r
}

func decodeCreateItemRequest(c context.Context, r *http.Request) (interface{}, error) {
	var body = representation.ItemEntity{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return endpoint.CreateItemRequest{Data: &body}, nil
}

func decodeShowItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}

	i, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return endpoint.ShowItemRequest{ID: i}, nil
}
