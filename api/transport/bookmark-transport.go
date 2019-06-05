package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/asxcandrew/galas/faults"
	"github.com/asxcandrew/galas/social/bookmark"

	"github.com/asxcandrew/galas/api/endpoint"
	"github.com/asxcandrew/galas/worker"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeBookmarkHandler(s bookmark.BookmarkService, w worker.AuthWorker, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(worker.HTTPToContext()),
	}

	listBookmarksHandler := kithttp.NewServer(
		w.NewJWTParser(endpoint.MakeListBookmarksEndpoint(s)),
		decodeListBookmarksRequest,
		encodeResponse,
		opts...,
	)

	createBookmarkHandler := kithttp.NewServer(
		w.NewJWTParser(endpoint.MakeCreateBookmarkEndpoint(s)),
		decodeCreateBookmarkRequest,
		encodeResponse,
		opts...,
	)

	deleteBookmarkHandler := kithttp.NewServer(
		w.NewJWTParser(endpoint.MakeDeleteBookmarkEndpoint(s)),
		decodeDeleteBookmarkRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/api/v1/bookmarks/", listBookmarksHandler).Methods("GET")
	r.Handle("/api/v1/bookmarks/", createBookmarkHandler).Methods("POST")
	r.Handle("/api/v1/bookmarks/{id}", deleteBookmarkHandler).Methods("DELETE")

	return r
}

func decodeCreateBookmarkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = endpoint.CreateBookmarkRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeDeleteBookmarkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	val, ok := vars["id"]

	if !ok {
		return nil, errors.New(faults.BadRequestError)
	}

	i, err := strconv.Atoi(val)

	if err != nil {
		return nil, err
	}

	return endpoint.DeleteBookmarkRequest{ID: i}, nil
}

func decodeListBookmarksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	page := getPage(r)

	return endpoint.ListBookmarksRequest{Page: page}, nil
}
