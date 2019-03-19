package user

import (
	"github.com/asxcandrew/galas/storage"
)

type UserService interface {
}

type userService struct {
	storage storage.Storage
}

// NewUserService creates an usre service with necessary dependencies.
func NewUserService(storage storage.Storage) UserService {
	return &userService{
		storage: storage,
	}
}
