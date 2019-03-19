package item

import (
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
)

type ItemService interface {
	Get(int) (*model.Item, error)
}

type itemService struct {
	storage storage.Storage
}

// NewItemService creates an item service with necessary dependencies.
func NewItemService(storage storage.Storage) ItemService {
	return &itemService{
		storage: storage,
	}
}

func (s *itemService) Get(itemID int) (*model.Item, error) {
	return &model.Item{}, nil
}
