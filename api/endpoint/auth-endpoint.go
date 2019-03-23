package endpoint

import (
	"context"
	"time"

	"github.com/asxcandrew/galas/storage/model"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/workers"
	"github.com/go-kit/kit/endpoint"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type AuthResponse struct {
	User   *representation.UserEntity `json:"user"`
	Token  string                     `json:"token"`
	Expire string                     `json:"expire"`
}

func MakeLoginEndpoint(w workers.AuthWorker) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		user, err := w.Login(req.Email, req.Password)

		if err != nil {
			return nil, nil
		}

		token, expire, err := w.GenerateToken(user)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = buildAuthResponse(user, token, expire)
		}
		return resp, err
	}
}

func MakeRegisterEndpoint(w workers.AuthWorker) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		user := &model.User{
			Email:    req.Email,
			Username: req.Username,
			Role:     model.UserRole_Plebs,
		}
		err := w.Register(user, req.Password)

		if err != nil {
			return nil, nil
		}

		token, expire, err := w.GenerateToken(user)

		resp := representation.Resp{
			Err: err,
		}
		if err == nil {
			resp.Data = buildAuthResponse(user, token, expire)
		}
		return resp, err
	}
}

func buildAuthResponse(user *model.User, token string, expire time.Time) *AuthResponse {
	return &AuthResponse{
		User: &representation.UserEntity{
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			About:     user.About,
		},
		Token:  token,
		Expire: expire.Format(time.RFC3339),
	}
}
