package endpoint

import (
	"context"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/bookmark"
	"github.com/asxcandrew/galas/errors"
	"github.com/asxcandrew/galas/workers"
	"github.com/go-kit/kit/endpoint"
)

type DeleteBookmarkRequest struct {
	ID int
}

type CreateBookmarkRequest struct {
	ItemID  int    `json:"item_id"`
	Comment string `json:"comment"`
}

type ListBookmarksRequest struct {
	Page int
}

func MakeListBookmarksEndpoint(s bookmark.BookmarkService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListBookmarksRequest)
		claims, err := workers.GetClaims(ctx)

		if err != nil {
			return nil, err
		}

		bookmarks, err := s.List(claims.UserID, req.Page)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = representation.ConvertBookmarksListModelToEntity(bookmarks)
		}
		return resp, err
	}
}

func MakeCreateBookmarkEndpoint(s bookmark.BookmarkService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		claims, err := workers.GetClaims(ctx)

		if err != nil {
			return nil, err
		}

		req := request.(CreateBookmarkRequest)
		bookmark, err := s.Create(req.ItemID, claims.UserID, req.Comment)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = representation.ConvertBookmarkModelToEntity(bookmark)
		}
		return resp, err
	}
}

func MakeDeleteBookmarkEndpoint(s bookmark.BookmarkService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		claims, err := workers.GetClaims(ctx)

		if err != nil {
			return nil, err
		}

		req := request.(DeleteBookmarkRequest)
		bookmark, err := s.GetByID(req.ID)

		if bookmark.UserID != claims.UserID {
			return nil, errors.ForbiddenError
		}
		err = s.Delete(req.ID)

		return nil, err
	}
}
