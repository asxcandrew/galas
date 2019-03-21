package repository

import (
	"github.com/asxcandrew/galas/storage/model"
)

type ItemRepository interface {
	GetByID(int) (*model.Item, error)
	Create(*model.Item) error
}
