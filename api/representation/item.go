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
	AuthorID   int       `json:"author_id"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
}

func ConvertItemModelToEntity(m *model.Item) *ItemEntity {
	return &ItemEntity{
		ID:        m.ID,
		Link:      m.Link,
		AuthorID:  m.AuthorID,
		HTMLBody:  m.HTMLBody,
		Title:     m.Title,
		Type:      m.Type,
		Score:     m.Score,
		CreatedAt: m.CreatedAt,
	}
}
