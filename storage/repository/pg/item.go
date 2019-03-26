package pg

import (
	"github.com/asxcandrew/galas/storage/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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
	item = &model.Item{}
	err = r.db.Model(item).Where("item.id = ?", id).Column("Author").Select()
	return item, err
}

func (r *ItemRepository) GetNewStories(page int) ([]*model.Item, error) {
	var items []*model.Item

	q := r.activeQuery(r.db.Model(&items))
	q = q.Where("type = ?", model.ItemType_Story)
	q, err := paginate(q, page)

	if err != nil {
		return items, err
	}

	err = q.Select()

	return items, err
}

func (r *ItemRepository) Create(item *model.Item) (err error) {
	return create(r.db, item)
}

func (r *ItemRepository) activeQuery(q *orm.Query) *orm.Query {
	return q.Where("active = ?", true).Column("Author").Order("created_at ASC")
}
