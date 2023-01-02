package logs

import (
	"errors"
	"fmt"
)

type Config struct {
	MinLevel uint8  `yaml:"min_level"`
	FilePath string `yaml:"file_path"`
}

func (conf *Config) Init() {
	conf.MinLevel = LevelInfo
	conf.FilePath = "/dev/stderr"
}

func (conf *Config) Validate() error {
	if conf.MinLevel < MinLogLevel || conf.MinLevel > MaxLogLevel {
		return fmt.Errorf(`invalid "min_level" [value = %d]`, conf.MinLevel)
	}
	if conf.FilePath == "" {
		return errors.New(`empty "file_path"`)
	}
	return nil
}

func (conf *Config) Prepare() {}
