package storage

import (
	"errors"
)

type Config struct {
	Database string `yaml:"database"`

	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (conf *Config) Init() {}

func (conf *Config) Validate() error {
	if conf.Database == "" {
		return errors.New(`no "database"`)
	}

	if conf.Host == "" {
		return errors.New(`no "host"`)
	}
	if conf.Port == 0 {
		return errors.New(`no "port"`)
	}

	if conf.Username == "" {
		return errors.New(`no "username"`)
	}
	if conf.Password == "" {
		return errors.New(`no "password"`)
	}

	return nil
}

func (conf *Config) Prepare() {}
