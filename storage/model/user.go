package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

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
	MediaID           int
	Media             *Media
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required, validation.Length(3, 10)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.EncryptedPassword, validation.Required),
		validation.Field(&u.Role,
			validation.Required,
			validation.In(
				UserRole_Patrician,
				UserRole_Plebs,
			),
		),
	)
}
