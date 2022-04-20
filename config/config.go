package config

import (
	"github.com/jfonseca85/aws-sdk-hello-world/log"
	"github.com/spf13/viper"
)

var (
	_ Config = (*config)(nil)
)

type Config interface {
	Load() (*config, error)
}

type config struct {
	v   *viper.Viper
	log *log.Helper
}

// New new a config with options.
func New(opts ...Option) Config {
	viper := viper.GetViper()
	for _, opt := range opts {
		opt.apply(viper)
	}
	return &config{
		v:   viper,
		log: log.NewHelper(log.GetLogger()),
	}
}

func (c *config) Load() (*config, error) {
	err := c.v.ReadInConfig()
	if err != nil {
		return &config{}, err
	}
	return c, nil

}
