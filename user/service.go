package user

import (
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
)

type UserService interface {
	Get(string) (*model.User, error)
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

func (s *userService) Get(username string) (user *model.User, err error) {
	user, err = s.storage.User.GetByUsername(username)

	return user, err
}
