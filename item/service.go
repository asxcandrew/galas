package item

import (
	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
)

type ItemService interface {
	Get(int) (*model.Item, error)
	ListNew(int) ([]*model.Item, error)
	Create(*representation.ItemEntity, int) (*model.Item, error)
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

func (s *itemService) ListNew(page int) ([]*model.Item, error) {
	st, e := s.storage.Item.GetNewStories(page)
	return st, e
}

func (s *itemService) Create(item *representation.ItemEntity, authorID int) (*model.Item, error) {
	i := &model.Item{
		Link:       item.Link,
		Score:      0,
		HTMLBody:   item.HTMLBody,
		Title:      item.Title,
		Type:       item.Type,
		AncestorID: item.AncestorID,
		AuthorID:   authorID,
		Active:     true,
	}
	err := i.Validate()

	if err != nil {
		return nil, err
	}

	err = s.storage.Item.Create(i)

	if err != nil {
		return nil, err
	}

	i, err = s.storage.Item.GetByID(i.ID)

	return i, err
}
