package endpoint

import (
	"context"
	"fmt"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/faults"
	"github.com/asxcandrew/galas/social/user"
	"github.com/asxcandrew/galas/storage/model"
	"github.com/asxcandrew/galas/worker"
	"github.com/go-kit/kit/endpoint"
)

type GetUserRequest struct {
	Username string
}

type UpdateUserRequest struct {
	Username string
	Data     struct {
		About   string `json:"about"`
		MediaID int    `json:"media_id"`
	}
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

func MakeUpdateUserEndpoint(s user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)

		claims, err := worker.GetClaims(ctx)

		if err != nil {
			return nil, err
		}
		if claims.UserName != req.Username {
			return nil, faults.BuildRichError(
				faults.ForbiddenError,
				fmt.Errorf("User: %s, forbidden to update: %s", claims.UserName, req.Username),
			)
		}

		user := &model.User{
			ID:      claims.UserID,
			About:   req.Data.About,
			MediaID: req.Data.MediaID,
		}

		err = s.Update(user)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = representation.ConvertUserModelToEntity(user)
		}
		return resp, err
	}
}
