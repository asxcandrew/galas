package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	err := migrations.Register(func(db migrations.DB) error {
		_, err := db.Exec(``)
		return err
	},
		func(db migrations.DB) error {
			_, err := db.Exec(``)
			return err
		})

	fmt.Println(err)
}
