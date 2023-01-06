package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/denis96z/simple-version-tracker/worker/pkg/checker"
	"github.com/denis96z/simple-version-tracker/worker/pkg/logs"
	"github.com/denis96z/simple-version-tracker/worker/pkg/loop"
	"github.com/denis96z/simple-version-tracker/worker/pkg/storage"
)

type Config struct {
	Loop    loop.Config    `yaml:"loop"`
	Logger  logs.Config    `yaml:"logger"`
	Storage storage.Config `yaml:"storage"`
	Checker checker.Config `yaml:"checker"`
}

func Load(confFilePath string) (*Config, error) {
	conf := &Config{}
	conf.Init()

	b, err := os.ReadFile(confFilePath)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read config file [path = %q]: %w",
			confFilePath, err,
		)
	}
	if err = yaml.Unmarshal(b, conf); err != nil {
		return nil, fmt.Errorf(
			"failed to parse yaml config file: %w", err,
		)
	}
	if err = conf.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	conf.Prepare()
	return conf, nil
}

func (conf *Config) Init() {
	conf.Loop.Init()
	conf.Logger.Init()
	conf.Storage.Init()
	conf.Checker.Init()
}

func (conf *Config) Validate() error {
	if err := conf.Loop.Validate(); err != nil {
		return fmt.Errorf(
			`"loop" config validation failed: %w`, err,
		)
	}
	if err := conf.Logger.Validate(); err != nil {
		return fmt.Errorf(
			`"logger" config validation failed: %w`, err,
		)
	}
	if err := conf.Storage.Validate(); err != nil {
		return fmt.Errorf(
			`"storage" config validation failed: %w`, err,
		)
	}
	if err := conf.Checker.Validate(); err != nil {
		return fmt.Errorf(
			`"checker" config validation failed: %w`, err,
		)
	}
	return nil
}

func (conf *Config) Prepare() {
	conf.Loop.Prepare()
	conf.Logger.Prepare()
	conf.Storage.Prepare()
	conf.Checker.Prepare()
}

func (conf *Config) Dump() string {
	b, err := yaml.Marshal(*conf)
	if err != nil {
		panic(err)
	}

	return string(b)
}
