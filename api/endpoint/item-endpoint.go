package endpoint

import (
	"context"

	"github.com/asxcandrew/galas/item"
	"github.com/go-kit/kit/endpoint"
)

type ShowItemRequest struct {
	ID int
}

type ShowItemResponse struct {
	ID  int   `json:"id,omitempty"`
	Err error `json:"error,omitempty"`
}

func (r ShowItemResponse) error() error { return r.Err }

func MakeShowItemEndpoint(s item.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ShowItemRequest)
		item, err := s.Get(req.ID)
		return ShowItemResponse{ID: item.ID, Err: err}, nil
	}
}
