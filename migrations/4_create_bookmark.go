package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table bookmarks...")
		_, err := db.Exec(`CREATE TABLE bookmarks (
												id serial PRIMARY KEY,
												comment varchar,
												user_id int REFERENCES users(id),
												item_id int REFERENCES items(id),
												updated_at timestamptz DEFAULT current_timestamp,
												created_at timestamptz DEFAULT current_timestamp,
												CONSTRAINT unq_user_id_item_id UNIQUE(user_id,item_id)
											)
		`)
		return err
	},
		func(db migrations.DB) error {
			fmt.Println("dropping table bookmarks...")
			_, err := db.Exec(`DROP TABLE bookmarks`)
			return err
		})
}
