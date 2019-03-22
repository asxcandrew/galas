package representation

import (
	"time"
)

type UserEntity struct {
	Username  string    `json:"username"`
	About     string    `json:"about"`
	Karma     int       `json:"karma"`
	Items     *[]int    `json:"items"`
	CreatedAt time.Time `json:"created_at"`
}
