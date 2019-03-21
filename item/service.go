package item

import (
	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
)

type ItemService interface {
	Get(int) (*model.Item, error)
	Create(*representation.ItemEntity) (*model.Item, error)
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

func (s *itemService) Get(itemID int) (item *model.Item, err error) {
	item, err = s.storage.Item.GetByID(itemID)

	return item, err
}

func (s *itemService) Create(item *representation.ItemEntity) (*model.Item, error) {
	model := &model.Item{
		Link:       item.Link,
		Score:      0,
		HTMLBody:   item.HTMLBody,
		Title:      item.Title,
		Type:       item.Type,
		AncestorID: item.AncestorID,
		AuthorID:   item.AuthorID,
		Active:     true,
	}
	err := model.Validate()

	if err != nil {
		return nil, err
	}

	err = s.storage.Item.Create(model)

	return model, err
}
