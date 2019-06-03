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
	var fieldRules []*validation.FieldRules

	fieldRules = append(fieldRules, validation.Field(&i.HTMLBody, validation.Required))
	fieldRules = append(fieldRules, validation.Field(&i.Type,
		validation.Required,
		validation.In(
			ItemType_Story,
			ItemType_Comment,
			ItemType_Announcement,
		),
	))

	if i.Type == ItemType_Story {
		fieldRules = append(fieldRules, validation.Field(&i.Title, validation.Required, validation.Length(5, 50)))
	}

	if i.Type == ItemType_Comment {
		fieldRules = append(fieldRules, validation.Field(&i.AncestorID, validation.Required))
	}

	return validation.ValidateStruct(i, fieldRules...)
}
