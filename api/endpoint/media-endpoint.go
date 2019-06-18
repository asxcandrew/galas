package endpoint

import (
	"context"
	"strconv"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/social/media"
	"github.com/go-kit/kit/endpoint"
)

type CreateMediaRequest struct {
	ContentType string `json:"content_type"`
}

func MakeGetMediaEndpoint(s media.MediaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}

func MakeCreateMediaEndpoint(s media.MediaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateMediaRequest)
		media, url, err := s.Create(req.ContentType)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = map[string]string{
				"url":      url.String(),
				"media_id": strconv.Itoa(media.ID),
			}
		}
		return resp, err
	}
}

func MakeDeleteMediaEndpoint(s media.MediaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}
