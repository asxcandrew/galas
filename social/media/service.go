package media

import (
	"net/url"

	"github.com/asxcandrew/galas/storage"
	"github.com/asxcandrew/galas/storage/model"
	"github.com/asxcandrew/galas/worker"
	uuid "github.com/satori/go.uuid"
)

type MediaService interface {
	Create(string) (*model.Media, *url.URL, error)
	Delete(string) error
}

type mediaService struct {
	storage       storage.Storage
	objectStorage worker.ObjectStorage
}

// NewMediaService creates an media service with necessary dependencies.
func NewMediaService(s storage.Storage, o worker.ObjectStorage) MediaService {
	return &mediaService{
		storage:       s,
		objectStorage: o,
	}
}

func (s *mediaService) Create(contentType string) (*model.Media, *url.URL, error) {
	objectName := uuid.NewV4()
	presignedURL, err := s.objectStorage.PresignedPutObject(objectName.String())

	if err != nil {
		return nil, nil, err
	}

	media := model.Media{
		ContentType: contentType,
		Name:        objectName.String(),
	}

	err = s.storage.Media.Create(&media)

	if err != nil {
		return nil, nil, err
	}

	return &media, presignedURL, nil
}

func (s *mediaService) Delete(mediaUUID string) error {
	err := s.storage.Media.Delete(mediaUUID)

	return err
}
