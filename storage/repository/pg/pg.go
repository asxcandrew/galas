package pg

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/go-pg/pg/urlvalues"
)

const (
	PerPage = 20
)

func newPager(page int) *urlvalues.Pager {
	p := urlvalues.Pager{
		Limit: PerPage,
	}

	p.SetPage(page)

	return &p
}

func create(db *pg.DB, model interface{}) error {
	return db.Insert(model)
}

func paginate(q *orm.Query, page int) (*orm.Query, error) {
	p := newPager(page)

	return p.Pagination(q)
}