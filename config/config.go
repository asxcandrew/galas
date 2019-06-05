package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type galasConfiguration struct {
	DB          *databaseConfiguration
	FileStorage *fileStorageConfiguration
	SecretSeed  string `required:"true"`
}

type databaseConfiguration struct {
	Host     string `required:"true"`
	Port     string `default:"5432"`
	Name     string `required:"true"`
	Password string `required:"true"`
	User     string `required:"true"`
}

type fileStorageConfiguration struct {
	Endpoint  string `required:"true"`
	Path      string `required:"true"`
	Bucket    string `required:"true"`
	AccessKey string `split_words:"true" required:"true"`
	SecretKey string `split_words:"true" required:"true"`
}

func ResolveConfig() (*galasConfiguration, error) {
	config := &galasConfiguration{}

	//
	// Resolve env. variables
	//
	if err := envconfig.Process("GALAS", config); err != nil {
		return nil, fmt.Errorf("failed to parse environment configurations, %s", err.Error())
	}

	return config, nil
}
