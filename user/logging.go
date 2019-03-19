package user

import (
	"github.com/go-kit/kit/log"
)

type userLoggingService struct {
	logger log.Logger
	UserService
}

// NewLoggingService returns a new instance of a itemLoggingService.
func NewUserLoggingService(logger log.Logger, s UserService) UserService {
	return &userLoggingService{logger, s}
}
