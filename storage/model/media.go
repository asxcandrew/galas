package model

import (
	"time"
)

//
// Media type reperesents the db structure of media.
//
type Media struct {
	ID          int
	ContentType string
	Name        string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
