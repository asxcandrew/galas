package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table media...")
		_, err := db.Exec(`CREATE TABLE media (
												id serial PRIMARY KEY,
												name varchar UNIQUE,
												content_type varchar,
												updated_at timestamptz DEFAULT current_timestamp,
												created_at timestamptz DEFAULT current_timestamp
											)
		`)
		return err
	},
		func(db migrations.DB) error {
			fmt.Println("dropping table media...")
			_, err := db.Exec(`DROP TABLE media`)
			return err
		})
}
