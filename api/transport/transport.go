package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/asxcandrew/galas/api/representation"
)

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
