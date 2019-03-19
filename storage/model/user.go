package model

import "time"

const (
	UserRole_Patrician = "patrician"
	UserRole_Plebs     = "plebs"
)

//
// User type reperesents the db structure of user.
//
type User struct {
	ID                int
	About             string
	Username          string
	Role              string
	Email             string
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
