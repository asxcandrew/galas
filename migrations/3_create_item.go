package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table items...")
		_, err := db.Exec(`CREATE TABLE items (
												id serial PRIMARY KEY,
												link varchar,
												html_body varchar,
												title varchar,
												type varchar,
												author_id int REFERENCES users(id),
												ancestor_id int REFERENCES items(id),
												score int,
												active boolean DEFAULT TRUE,
												updated_at timestamptz DEFAULT current_timestamp,
												created_at timestamptz DEFAULT current_timestamp
											)
		`)
		return err
	},
		func(db migrations.DB) error {
			fmt.Println("dropping table items...")
			_, err := db.Exec(`DROP TABLE items`)
			return err
		})
}
