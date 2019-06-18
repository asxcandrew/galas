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
	err := create(r.db, b)
	return wrapError(err)
}

func (r *BookmarkRepository) ListAndCount(userID, page int) ([]*model.Bookmark, int, error) {
	var bookmarks []*model.Bookmark

	q := r.db.Model(&bookmarks)
	q.Where("user_id = ?", userID).Order("id DESC").Column("Item", "Item.Author")

	q, err := paginate(q, page)

	if err != nil {
		return bookmarks, 0, wrapError(err)
	}

	count, err := q.SelectAndCount()

	return bookmarks, count, wrapError(err)
}

func (r *BookmarkRepository) Delete(ID int) error {
	b := &model.Bookmark{ID: ID}
	err := r.db.Delete(b)

	return wrapError(err)
}

func (r *BookmarkRepository) GetByItemID(ItemID int) (*model.Bookmark, error) {
	b := &model.Bookmark{}
	err := r.db.Model(b).Where("bookmark.item_id = ?", ItemID).Column("Item").Select()

	return b, wrapError(err)
}

func (r *BookmarkRepository) PerPage() int {
	return perPage
}
