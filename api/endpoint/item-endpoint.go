package endpoint

import (
	"context"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/item"
	"github.com/asxcandrew/galas/workers"
	"github.com/go-kit/kit/endpoint"
)

type ShowItemRequest struct {
	ID int
}

type CreateItemRequest struct {
	Data *representation.ItemEntity
}

func MakeShowItemEndpoint(s item.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ShowItemRequest)
		item, err := s.Get(req.ID)
		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = representation.ConvertItemModelToEntity(item)
		}
		return resp, err
	}
}

func MakeCreateItemEndpoint(s item.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		claims, err := workers.GetClaims(ctx)

		if err != nil {
			return nil, err
		}

		req := request.(CreateItemRequest)
		req.Data.AuthorID = claims.UserID
		item, err := s.Create(req.Data)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = representation.ConvertItemModelToEntity(item)
		}
		return resp, err
	}
}
