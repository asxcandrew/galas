package model

import (
	"time"
)

//
// Bookmark type reperesents the db structure of user.
//
type Bookmark struct {
	ID        int
	UserID    int
	User      *User
	ItemID    int
	Item      *Item
	Comment   string
	CreatedAt time.Time
}
