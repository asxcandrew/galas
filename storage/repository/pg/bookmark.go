package pg

import (
	"github.com/asxcandrew/galas/storage/model"
	"github.com/go-pg/pg"
)

type BookmarkRepository struct {
	db *pg.DB
}

func NewPGBookmarkRepository(db *pg.DB) *BookmarkRepository {
	return &BookmarkRepository{
		db: db,
	}
}

func (r *BookmarkRepository) Create(b *model.Bookmark) error {
	return create(r.db, b)
}

func (r *BookmarkRepository) List(userID, page int) ([]*model.Bookmark, error) {
	var bookmarks []*model.Bookmark

	q := r.db.Model(&bookmarks)
	q.Where("user_id = ?", userID).Column("Item")

	q, err := paginate(q, page)

	if err != nil {
		return bookmarks, err
	}

	return bookmarks, err
}

func (r *BookmarkRepository) Delete(ID int) error {
	b := &model.Bookmark{ID: ID}
	err := r.db.Delete(b)

	return wrapError(err)
}

func (r *BookmarkRepository) GetByID(ID int) (*model.Bookmark, error) {
	b := &model.Bookmark{ID: ID}
	err := r.db.Select(b)

	return b, wrapError(err)
}