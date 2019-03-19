package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type galasConfiguration struct {
	DB *databaseConfiguration
}

type databaseConfiguration struct {
	Host     string `required:"true"`
	Port     string `default:"5432"`
	Name     string `required:"true"`
	Password string `required:"true"`
	User     string `required:"true"`
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
