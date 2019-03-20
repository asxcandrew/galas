package storage

import (
	"fmt"

	"github.com/go-pg/pg"
)

func InitPGConnection(host, port, username, password, name string) (db *pg.DB) {
	options := pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		User:     username,
		Password: password,
		Database: name,
	}
	// if c.TLS {
	// 	options.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// }
	db = pg.Connect(&options)
	return db
}
