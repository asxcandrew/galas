package representation

import (
	"time"

	"github.com/asxcandrew/galas/storage/model"
)

type UserEntity struct {
	Username  string    `json:"username"`
	About     string    `json:"about"`
	Karma     int       `json:"karma"`
	Items     *[]int    `json:"items"`
	CreatedAt time.Time `json:"created_at"`
}

func ConvertUserModelToEntity(m *model.User) *UserEntity {
	return &UserEntity{
		Username:  m.Username,
		About:     m.About,
		CreatedAt: m.CreatedAt,
	}
}
