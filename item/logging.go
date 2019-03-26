package item

import (
	"fmt"
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
	logger = log.With(logger, "service", "items")

	return &itemLoggingService{logger, s}
}

func (s *itemLoggingService) Get(id int) (item *model.Item, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get",
			"params", fmt.Sprintf("[id=%d]", id),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.ItemService.Get(id)
}

func (s *itemLoggingService) Create(item *representation.ItemEntity, authorID int) (res *model.Item, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.ItemService.Create(item, authorID)
}

func (s *itemLoggingService) ListNew(page int) (res []*model.Item, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "feed_new",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.ItemService.ListNew(page)
}
