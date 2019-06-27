package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asxcandrew/galas/worker"

	"github.com/asxcandrew/galas/api/transport"
	"github.com/asxcandrew/galas/config"
	"github.com/asxcandrew/galas/social/bookmark"
	"github.com/asxcandrew/galas/social/item"
	"github.com/asxcandrew/galas/social/media"
	"github.com/asxcandrew/galas/social/user"
	"github.com/asxcandrew/galas/storage"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var logger log.Logger

func main() {
	var wait time.Duration

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	var httpAddr = ":8000"

	errs := make(chan error, 1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "galas", log.DefaultTimestampUTC)

	appConfig, err := config.ResolveConfig()

	if err != nil {
		errs <- err
	}

	db := storage.InitPGConnection(
		appConfig.DB.Host,
		appConfig.DB.Port,
		appConfig.DB.User,
		appConfig.DB.Password,
		appConfig.DB.Name,
	)

	defer db.Close()

	objStorage, err := worker.NewObjectStorage(
		&worker.ObjectStorageConfig{
			Endpoint:  appConfig.FileStorage.Endpoint,
			Bucket:    appConfig.FileStorage.Bucket,
			Path:      appConfig.FileStorage.Path,
			AccessKey: appConfig.FileStorage.AccessKey,
			SecretKey: appConfig.FileStorage.SecretKey,
		},
	)

	if err != nil {
		errs <- err
	}

	st := storage.NewPGStorage(db)
	us := user.NewUserService(st)
	is := item.NewItemService(st)
	bs := bookmark.NewBookmarkService(st)
	ms := media.NewMediaService(st, objStorage)

	is = item.NewItemLoggingService(logger, is)
	us = user.NewUserLoggingService(logger, us)
	bs = bookmark.NewBookmarkLoggingService(logger, bs)

	aw := worker.NewAuthWorker(appConfig.SecretSeed)

	httpLogger := log.With(logger, "component", "http")

	routes := mux.NewRouter()
	api := routes.PathPrefix("/api/v1").Subrouter()

	api.PathPrefix("/users").Handler(transport.MakeUserHandler(us, aw, httpLogger))
	api.PathPrefix("/items").Handler(transport.MakeItemHandler(is, aw, httpLogger))
	api.PathPrefix("/bookmarks").Handler(transport.MakeBookmarkHandler(bs, aw, httpLogger))
	api.PathPrefix("/media").Handler(transport.MakeMediaHandler(ms, httpLogger))
	api.PathPrefix("/feed").Handler(transport.MakeFeedHandler(is, aw, httpLogger))
	api.PathPrefix("/auth").Handler(transport.MakeAuthHandler(us, aw, httpLogger))

	srv := &http.Server{
		Addr:         httpAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes,
	}

	go func() {
		logger.Log("transport", "http", "address", httpAddr, "msg", "listening")

		errs <- srv.ListenAndServe()
	}()

	select {
	case <-c:
		shutdown(srv, wait)
	case <-errs:
		shutdown(srv, wait)
	}
}

func shutdown(srv *http.Server, wait time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Log("transport", "shutting down...")
	os.Exit(0)
}
