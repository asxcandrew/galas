package pg

import (
	"github.com/go-pg/pg"
)

type UserRepository struct {
	db *pg.DB
}

func NewPGUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
