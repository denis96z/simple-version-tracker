package checker

import (
	"errors"
	"fmt"
)

type Config struct {
	Docker DockerConfig `yaml:"docker"`
}

func (conf *Config) Init() {
	conf.Docker.Init()
}

func (conf *Config) Validate() error {
	if err := conf.Docker.Validate(); err != nil {
		return fmt.Errorf(
			`"docker" config validation failed: %w`, err,
		)
	}
	return nil
}

func (conf *Config) Prepare() {
	conf.Docker.Prepare()
}

type DockerConfig struct {
	BinaryPath          string `yaml:"binary_path"`
	ConfigDirectoryPath string `yaml:"config_directory_path"`
}

func (conf *DockerConfig) Init() {}

func (conf *DockerConfig) Validate() error {
	if conf.BinaryPath == "" {
		return errors.New(`empty "binary_path"`)
	}
	if conf.ConfigDirectoryPath == "" {
		return errors.New(`empty "config_directory_path"`)
	}
	return nil
}

func (conf *DockerConfig) Prepare() {}
