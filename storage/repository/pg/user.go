package pg

import (
	"github.com/asxcandrew/galas/storage/model"
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

func (r *UserRepository) Create(user *model.User) error {
	return create(r.db, user)
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Model(user).Where("email = ?", email).Select()

	return user, err
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Model(user).Where("username = ?", username).Select()

	return user, wrapError(err)
}
