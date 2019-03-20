package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`CREATE TABLE users (
												id serial PRIMARY KEY,
												username varchar UNIQUE,
												email varchar UNIQUE,
												about varchar,
												role varchar,
												encrypted_password varchar,
												updated_at timestamptz DEFAULT current_timestamp,
												created_at timestamptz DEFAULT current_timestamp
											)
		`)
		return err
	},
		func(db migrations.DB) error {
			fmt.Println("dropping table users...")
			_, err := db.Exec(`DROP TABLE users`)
			return err
		})
}
