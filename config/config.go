package config

import (
	"github.com/jfonseca85/aws-sdk-hello-world/log"
	"github.com/spf13/viper"
)

var _ Config = (*config)(nil)

type Config interface {
	Load() (*config, error)
}

type config struct {
	v   *viper.Viper
	log *log.Helper
}

// New new a config with options.
func New(opts ...Option) Config {
	o := options{
		v:      viper.GetViper(),
		logger: log.GetLogger(),
	}

	for _, opt := range opts {
		opt(&o)
	}
	return &config{
		v:   o.v,
		log: log.NewHelper(o.logger),
	}
}

func (c *config) Load() (*config, error) {
	err := c.v.ReadInConfig()
	if err != nil {
		return &config{}, err
	}
	return c, nil

}
