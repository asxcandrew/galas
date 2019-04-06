package representation

import (
	"time"

	"github.com/asxcandrew/galas/storage/model"
)

type BookmarkEntity struct {
	ID        int         `json:"id"`
	Comment   string      `json:"comment"`
	Item      *ItemEntity `json:"item"`
	CreatedAt time.Time   `json:"created_at"`
}

func ConvertBookmarkModelToEntity(m *model.Bookmark) *BookmarkEntity {
	return &BookmarkEntity{
		ID:        m.ID,
		Comment:   m.Comment,
		Item:      ConvertItemModelToEntity(m.Item),
		CreatedAt: m.CreatedAt,
	}
}

func ConvertBookmarksListModelToEntity(bookmarks []*model.Bookmark) []*BookmarkEntity {
	list := make([]*BookmarkEntity, len(bookmarks))

	for i, b := range bookmarks {
		list[i] = ConvertBookmarkModelToEntity(b)
	}
	return list
}
