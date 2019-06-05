package endpoint

import (
	"context"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/social/user"
	"github.com/go-kit/kit/endpoint"
)

type GetUserRequest struct {
	Username string
}

func MakeGetUserEndpoint(s user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := s.Get(req.Username)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = representation.ConvertUserModelToEntity(user)
		}
		return resp, err
	}
}
