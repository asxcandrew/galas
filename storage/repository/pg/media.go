package pg

import (
	"github.com/asxcandrew/galas/storage/model"
	"github.com/go-pg/pg"
)

type MediaRepository struct {
	db *pg.DB
}

func NewPGMediaRepository(db *pg.DB) *MediaRepository {
	return &MediaRepository{
		db: db,
	}
}

func (r *MediaRepository) Create(m *model.Media) error {
	return create(r.db, m)
}

func (r *MediaRepository) GetByID(id int) (*model.Media, error) {
	m := &model.Media{ID: id}
	err := r.db.Select(m)

	return m, err
}
