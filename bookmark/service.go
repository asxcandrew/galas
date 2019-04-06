package bookmark

import (
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
)

type BookmarkService interface {
	List(int, int) ([]*model.Bookmark, error)
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

func (s *bookmarkService) List(userID, page int) ([]*model.Bookmark, error) {
	bookmarks, err := s.storage.Bookmark.List(userID, page)

	return bookmarks, err
}

func (s *bookmarkService) Create(itemID int, userID int, comment string) (*model.Bookmark, error) {
	bookmark := &model.Bookmark{
		ItemID:  itemID,
		UserID:  userID,
		Comment: comment,
	}
	err := s.storage.Bookmark.Create(bookmark)

	bookmark, err = s.storage.Bookmark.GetByID(bookmark.ID)

	return bookmark, err
}

func (s *bookmarkService) Delete(bookmarkID int) error {
	err := s.storage.Bookmark.Delete(bookmarkID)

	return err
}

func (s *bookmarkService) GetByID(bookmarkID int) (*model.Bookmark, error) {
	b, err := s.storage.Bookmark.GetByID(bookmarkID)

	return b, err
}
