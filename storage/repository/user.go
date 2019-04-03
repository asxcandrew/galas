package repository

import (
	"github.com/asxcandrew/galas/storage/model"
)

type UserRepository interface {
	Create(*model.User) error
	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
}
