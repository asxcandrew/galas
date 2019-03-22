package pg

import (
	"github.com/asxcandrew/galas/storage/model"
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

func (r *ItemRepository) GetByID(id int) (item *model.Item, err error) {
	item = &model.Item{ID: id}
	err = r.db.Select(item)

	return item, err
}

func (r *ItemRepository) Create(item *model.Item) (err error) {
	return create(r.db, item)
}
