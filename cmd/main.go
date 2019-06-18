package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/asxcandrew/galas/worker"

	"github.com/asxcandrew/galas/api/transport"
	"github.com/asxcandrew/galas/config"
	"github.com/asxcandrew/galas/social/bookmark"
	"github.com/asxcandrew/galas/social/item"
	"github.com/asxcandrew/galas/social/media"
	"github.com/asxcandrew/galas/social/user"
	"github.com/asxcandrew/galas/storage"
	"github.com/go-kit/kit/log"
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

	db := storage.InitPGConnection(
		appConfig.DB.Host,
		appConfig.DB.Port,
		appConfig.DB.User,
		appConfig.DB.Password,
		appConfig.DB.Name,
	)

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
	mux := http.NewServeMux()

	mux.Handle("/api/v1/users/", transport.MakeUserHandler(us, httpLogger))
	mux.Handle("/api/v1/items/", transport.MakeItemHandler(is, aw, httpLogger))
	mux.Handle("/api/v1/bookmarks/", transport.MakeBookmarkHandler(bs, aw, httpLogger))
	mux.Handle("/api/v1/media/", transport.MakeMediaHandler(ms, httpLogger))
	mux.Handle("/api/v1/feed/", transport.MakeFeedHandler(is, aw, httpLogger))
	mux.Handle("/api/v1/auth/", transport.MakeAuthHandler(us, aw, httpLogger))

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
