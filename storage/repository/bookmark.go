package repository

import (
	"github.com/asxcandrew/galas/storage/model"
)

type BookmarkRepository interface {
	ListAndCount(userID, page int) ([]*model.Bookmark, int, error)
	Delete(int) error
	GetByItemID(int) (*model.Bookmark, error)
	Create(*model.Bookmark) error
	PerPage() int
}
