package representation

import (
	"time"

	"github.com/asxcandrew/galas/storage/model"
)

type ItemEntity struct {
	ID         int       `json:"id"`
	Link       string    `json:"link"`
	HTMLBody   string    `json:"html_body"`
	Title      string    `json:"title"`
	Type       string    `json:"type"`
	Score      int       `json:"score"`
	AncestorID int       `json:"ancestor_id"`
	Author     string    `json:"author"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
}

func ConvertItemModelToEntity(m *model.Item) *ItemEntity {
	return &ItemEntity{
		ID:         m.ID,
		Link:       m.Link,
		Author:     m.Author.Username,
		HTMLBody:   m.HTMLBody,
		AncestorID: m.AncestorID,
		Title:      m.Title,
		Type:       m.Type,
		Score:      m.Score,
		Active:     m.Active,
		CreatedAt:  m.CreatedAt,
	}
}

func ConvertItemsListModelToEntity(items []*model.Item) []*ItemEntity {
	list := make([]*ItemEntity, len(items))

	for i, item := range items {
		list[i] = ConvertItemModelToEntity(item)
	}
	return list
}
