package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/asxcandrew/galas/api/transport"
	"github.com/asxcandrew/galas/config"
	"github.com/asxcandrew/galas/item"
	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/user"
	"github.com/go-kit/kit/log"
	"github.com/go-pg/pg"
)

func main() {
	var logger log.Logger
	var httpAddr = ":8000"

	errs := make(chan error, 1)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "galas", log.DefaultTimestampUTC)

	appConfig, err := config.ResolveConfig()

	if err != nil {
		errs <- err
	}

	db := initPGConnection(
		appConfig.DB.Host,
		appConfig.DB.Port,
		appConfig.DB.User,
		appConfig.DB.Password,
		appConfig.DB.Name,
	)

	st := storage.NewPGStorage(db)
	us := user.NewUserService(st)
	is := item.NewItemService(st)

	is = item.NewItemLoggingService(logger, is)
	us = user.NewUserLoggingService(logger, us)

	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()

	mux.Handle("/api/v1/p/user/", transport.MakeUserHandler(us, httpLogger))
	mux.Handle("/api/v1/p/item/", transport.MakeItemHandler(is, httpLogger))

	http.Handle("/", mux)

	go func() {
		logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(httpAddr, mux)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)

	defer func() {
		db.Close()
	}()
}

func initPGConnection(host, port, username, password, name string) (db *pg.DB) {
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
