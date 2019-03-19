package storage

import (
	"github.com/asxcandrew/galas/storage/repository"
	pgrepository "github.com/asxcandrew/galas/storage/repository/pg"
	"github.com/go-pg/pg"
)

type Storage struct {
	Item repository.ItemRepository
	User repository.UserRepository
}

func NewPGStorage(db *pg.DB) Storage {
	return Storage{
		Item: pgrepository.NewPGItemRepository(db),
		User: pgrepository.NewPGUserRepository(db),
	}
}
