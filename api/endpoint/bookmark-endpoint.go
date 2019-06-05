package endpoint

import (
	"context"
	"errors"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/faults"
	"github.com/asxcandrew/galas/social/bookmark"
	"github.com/asxcandrew/galas/worker"
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

type ListBookmarkResponse struct {
	Bookmarks []*representation.BookmarkEntity `json:"bookmarks"`
	Total     int                              `json:"total"`
}

func MakeListBookmarksEndpoint(s bookmark.BookmarkService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListBookmarksRequest)
		claims, err := worker.GetClaims(ctx)

		if err != nil {
			return nil, err
		}

		bookmarks, count, err := s.List(claims.UserID, req.Page)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = &ListBookmarkResponse{
				Bookmarks: representation.ConvertBookmarksListModelToEntity(bookmarks),
				Total:     count,
			}
		}
		return resp, err
	}
}

func MakeCreateBookmarkEndpoint(s bookmark.BookmarkService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		claims, err := worker.GetClaims(ctx)

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

		claims, err := worker.GetClaims(ctx)

		if err != nil {
			return nil, err
		}

		req := request.(DeleteBookmarkRequest)
		bookmark, err := s.GetByID(req.ID)

		if err != nil {
			return nil, err
		}

		if bookmark.UserID != claims.UserID {
			return nil, faults.BuildRichError(faults.ForbiddenError, errors.New("User is forbidden to delete bookmark"))
		}
		err = s.Delete(req.ID)

		resp := representation.Resp{
			Err: err,
		}

		return resp, err
	}
}
