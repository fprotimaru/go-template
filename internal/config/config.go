package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PostgresUrl string `yaml:"postgresUrl"`
	SecretKey   string `yaml:"secretKey"`
}

func New(cfgFile string) (*Config, error) {
	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
