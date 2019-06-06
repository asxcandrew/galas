package worker

import (
	"errors"
	"net/url"
	"time"

	minio "github.com/minio/minio-go"
)

type objectStorage struct {
	client *minio.Client
	bucket string
	path   string
}

type ObjectStorageConfig struct {
	Endpoint  string
	Path      string
	Bucket    string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

type ObjectStorage interface {
	PresignedPutObject(string) (*url.URL, error)
	RemoveObject(string) error
	GetUrlForObject(string) (*url.URL, error)
}

func NewObjectStorage(c *ObjectStorageConfig) (ObjectStorage, error) {
	minioClient, err := minio.New(
		c.Endpoint,
		c.AccessKey,
		c.SecretKey,
		c.UseSSL,
	)

	if err != nil {
		return nil, err
	}

	exists, err := minioClient.BucketExists(c.Bucket)

	if err != nil || !exists {
		return nil, errors.New("Object storage: Bucket does not exist")
	}
	return &objectStorage{
		client: minioClient,
		bucket: c.Bucket,
		path:   c.Path,
	}, nil
}

func (s *objectStorage) PresignedPutObject(objectName string) (*url.URL, error) {
	url, err := s.client.PresignedPutObject(s.bucket, s.preparedName(objectName), time.Minute*5)
	return url, err
}

func (s *objectStorage) RemoveObject(objectName string) error {
	err := s.client.RemoveObject(s.bucket, objectName)
	return err
}

func (s *objectStorage) GetUrlForObject(objectName string) (*url.URL, error) {
	link, err := s.client.PresignedGetObject(s.bucket, s.preparedName(objectName), time.Hour, url.Values{})
	return link, err
}

func (s *objectStorage) preparedName(objectName string) string {
	return s.path + "/" + objectName
}
