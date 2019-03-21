package item

import (
	"time"

	"github.com/asxcandrew/galas/api/representation"
	"github.com/asxcandrew/galas/storage/model"
	"github.com/go-kit/kit/log"
)

type itemLoggingService struct {
	logger log.Logger
	ItemService
}

// NewItemLoggingService returns a new instance of a itemLoggingService.
func NewItemLoggingService(logger log.Logger, s ItemService) ItemService {
	return &itemLoggingService{logger, s}
}

func (s *itemLoggingService) Get(itemID int) (item *model.Item, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get",
			"id", itemID,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.ItemService.Get(itemID)
}

func (s *itemLoggingService) Create(item *representation.ItemEntity) (res *model.Item, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.ItemService.Create(item)
}
