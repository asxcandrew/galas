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
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var errBadRoute = errors.New("bad route")

func MakeItemHandler(s item.ItemService, logger log.Logger) http.Handler {
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
		endpoint.MakeCreateItemEndpoint(s),
		decodeCreateItemRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/api/v1/p/item/{id}", showItemHandler).Methods("GET")
	r.Handle("/api/v1/p/item/", createItemHandler).Methods("POST")

	return r
}

func decodeCreateItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if r, ok := response.(representation.Resp); ok {
		if r.GetError() != nil {
			encodeError(ctx, r.GetError(), w)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		return json.NewEncoder(w).Encode(map[string]interface{}{
			"data": r.GetData(),
		})
	}
	return errors.New("Bad request")
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
