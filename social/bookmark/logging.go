package bookmark

import (
	"github.com/go-kit/kit/log"
)

type bookmarkLoggingService struct {
	logger log.Logger
	BookmarkService
}

// NewBookmarkLoggingService returns a new instance of a itemLoggingService.
func NewBookmarkLoggingService(logger log.Logger, s BookmarkService) BookmarkService {
	return &bookmarkLoggingService{logger, s}
}
