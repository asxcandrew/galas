package repository

import (
	"github.com/asxcandrew/galas/storage/model"
)

type BookmarkRepository interface {
	ListAndCount(userID, page int) ([]*model.Bookmark, int, error)
	Delete(int) error
	GetByID(int) (*model.Bookmark, error)
	Create(*model.Bookmark) error
}
