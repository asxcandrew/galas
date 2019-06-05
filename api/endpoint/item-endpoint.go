package endpoint

import (
	"context"
	"errors"

	"github.com/asxcandrew/galas/storage/model"
	"github.com/asxcandrew/galas/worker"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/social/item"
	"github.com/go-kit/kit/endpoint"
)

const (
	ListType_New = "new"
	ListType_Top = "top"
)

type ShowItemRequest struct {
	ID int
}

type FeedRequest struct {
	Type string
	Page int
}

type CreateItemRequest struct {
	Data *representation.ItemEntity
}

func MakeFeedEndpoint(s item.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FeedRequest)
		var items []*model.Item
		var err error

		switch req.Type {
		case ListType_New:
			items, err = s.ListNew(req.Page)
		case ListType_Top:
			items, err = s.ListTop(req.Page)
		default:
			err = errors.New("Bad request")
		}

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = representation.ConvertItemsListModelToEntity(items)
		}
		return resp, err
	}
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

		claims, err := worker.GetClaims(ctx)

		if err != nil {
			return nil, err
		}

		req := request.(CreateItemRequest)
		item, err := s.Create(req.Data, claims.UserID)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = representation.ConvertItemModelToEntity(item)
		}
		return resp, err
	}
}
