package pg

import "github.com/go-pg/pg"

func create(db *pg.DB, model interface{}) error {
	return db.Insert(model)
}
