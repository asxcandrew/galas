package repository

import (
	"github.com/asxcandrew/galas/storage/model"
)

type MediaRepository interface {
	GetByID(int) (*model.Media, error)
	Create(*model.Media) error
}
