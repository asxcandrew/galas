package repository

import (
	"github.com/asxcandrew/galas/storage/model"
)

type ItemRepository interface {
	GetByID(int) (*model.Item, error)
	GetNewStories(int) ([]*model.Item, error)
	GetTopStories(int) ([]*model.Item, error)
	Create(*model.Item) error
}
