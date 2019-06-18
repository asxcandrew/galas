package media

import (
	"net/url"
	"time"

	"github.com/asxcandrew/galas/storage/model"
	"github.com/go-kit/kit/log"
)

type mediaLoggingService struct {
	logger log.Logger
	MediaService
}

// NewMediaLoggingService returns a new instance of a mediaLoggingService.
func NewMediaLoggingService(logger log.Logger, s MediaService) MediaService {
	logger = log.With(logger, "service", "media")

	return &mediaLoggingService{logger, s}
}

func (s *mediaLoggingService) Create(contentType string) (m *model.Media, u *url.URL, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.MediaService.Create(contentType)
}
