package pg

import (
	"github.com/go-pg/pg"
)

type ItemRepository struct {
	db *pg.DB
}

func NewPGItemRepository(db *pg.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}
