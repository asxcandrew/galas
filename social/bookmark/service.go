package bookmark

import (
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
)

type BookmarkService interface {
	List(int, int) ([]*model.Bookmark, int, error)
	Create(int, int, string) (*model.Bookmark, error)
	Delete(int) error
	GetByID(int) (*model.Bookmark, error)
}

type bookmarkService struct {
	storage storage.Storage
}

// NewBookmarkService creates an usre service with necessary dependencies.
func NewBookmarkService(storage storage.Storage) BookmarkService {
	return &bookmarkService{
		storage: storage,
	}
}

func (s *bookmarkService) List(userID, page int) ([]*model.Bookmark, int, error) {
	bookmarks, count, err := s.storage.Bookmark.ListAndCount(userID, page)

	return bookmarks, count, err
}

func (s *bookmarkService) Create(itemID int, userID int, comment string) (*model.Bookmark, error) {
	bookmark := &model.Bookmark{
		ItemID:  itemID,
		UserID:  userID,
		Comment: comment,
	}
	err := s.storage.Bookmark.Create(bookmark)

	if err != nil {
		return nil, err
	}

	item, err := s.storage.Item.GetByID(itemID)

	if err != nil {
		return nil, err
	}

	bookmark.Item = item

	return bookmark, nil
}

func (s *bookmarkService) Delete(bookmarkID int) error {
	err := s.storage.Bookmark.Delete(bookmarkID)

	return err
}

func (s *bookmarkService) GetByID(bookmarkID int) (*model.Bookmark, error) {
	b, err := s.storage.Bookmark.GetByID(bookmarkID)

	return b, err
}
