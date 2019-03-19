package model

import "time"

const (
	ItemType_story        = "story"
	ItemType_comment      = "comment"
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
