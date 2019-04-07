package user

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/asxcandrew/galas/storage/model"
)

type userLoggingService struct {
	logger log.Logger
	UserService
}

// NewLoggingService returns a new instance of a itemLoggingService.
func NewUserLoggingService(logger log.Logger, s UserService) UserService {
	logger = log.With(logger, "service", "users")

	return &userLoggingService{logger, s}
}

func (s *userLoggingService) Login(email, password string) (user *model.User, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "login",
			"email", email,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.UserService.Login(email, password)
}

func (s *userLoggingService) Register(user *model.User, password string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "register",
			"email", user.Email,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.UserService.Register(user, password)
}
