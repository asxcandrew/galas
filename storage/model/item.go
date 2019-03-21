package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	ItemType_Story        = "story"
	ItemType_Comment      = "comment"
	ItemType_Announcement = "announcement"
)

//
// Item type reperesents the db structure of Item.
//
type Item struct {
	ID         int
	Link       string
	HTMLBody   string
	Title      string
	Type       string
	Score      int
	AncestorID int
	AuthorID   int
	Author     *User
	Active     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (i *Item) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.Title, validation.Required, validation.Length(5, 50)),
		validation.Field(&i.HTMLBody, validation.Required),
		validation.Field(&i.Type,
			validation.Required,
			validation.In(
				ItemType_Story,
				ItemType_Comment,
				ItemType_Announcement,
			),
		),
	)
}
