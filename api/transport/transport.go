package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asxcandrew/galas/errors"

	"github.com/asxcandrew/galas/api/representation"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if r, ok := response.(representation.Resp); ok {
		if r.GetError() != nil {
			encodeError(ctx, r.GetError(), w)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"payload": r.GetData(),
		})
	}
	return errors.BadRequestError
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case errors.NotFoundError:
		w.WriteHeader(http.StatusNotFound)
	case errors.ForbiddenError:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func getPage(r *http.Request) int {
	p, ok := r.URL.Query()["page"]

	if !ok {
		return 0
	}

	if len(p[0]) > 0 {
		page, err := strconv.Atoi(p[0])

		if err != nil {
			return 0
		}
		return page
	}

	return 0
}
