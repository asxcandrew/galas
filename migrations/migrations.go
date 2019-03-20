package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/asxcandrew/galas/config"
	"github.com/asxcandrew/galas/storage"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

const UsageText = `This program runs command on the db. Supported commands are:
  - up - runs all available migrations.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
Usage:
  <command>
`

func main() {
	var db *pg.DB
	fmt.Printf(UsageText)

	flag.Parse()

	appConfig, err := config.ResolveConfig()

	if err != nil {
		exitf(err.Error())
	}

	db = storage.InitPGConnection(
		appConfig.DB.Host,
		appConfig.DB.Port,
		appConfig.DB.User,
		appConfig.DB.Password,
		appConfig.DB.Name,
	)

	fmt.Println(migrations.RegisteredMigrations())

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
	os.Exit(1)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
