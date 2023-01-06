package loop

import (
	"time"
)

type Config struct {
	Sleep time.Duration `yaml:"sleep"`
}

func (conf *Config) Init() {
	conf.Sleep = time.Minute
}

func (conf *Config) Validate() error {
	return nil
}

func (conf *Config) Prepare() {}
