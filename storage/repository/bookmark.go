package repository

import (
	"github.com/asxcandrew/galas/storage/model"
)

type BookmarkRepository interface {
	List(userID, page int) ([]*model.Bookmark, error)
	Delete(int) error
	GetByID(int) (*model.Bookmark, error)
	Create(*model.Bookmark) error
}
