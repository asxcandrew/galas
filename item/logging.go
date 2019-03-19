package item

import (
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
