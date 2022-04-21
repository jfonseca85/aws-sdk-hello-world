package config

import (
	"github.com/jfonseca85/aws-sdk-hello-world/log"
	"github.com/spf13/viper"
)

// Option is config option.
type Option func(*options)

type options struct {
	v      *viper.Viper
	logger log.Logger
}

// WithSource with config source.
func WithSource(s string) Option {
	return func(o *options) {
		o.v.SetConfigFile(s)
	}
}

//WithPath with config source.
func WithPath(s ...string) Option {
	return func(o *options) {
		for _, src := range s {
			o.v.AddConfigPath(src)
		}
	}
}

// WithPath with config source.
func WithName(s string) Option {
	return func(o *options) {
		o.v.SetConfigName(s)
	}
}

//WithType with config source.
func WithType(s string) Option {
	return func(o *options) {
		o.v.SetConfigType(s)
	}
}

// WithLogger with config logger.
func WithLogger(l log.Logger) Option {
	return func(o *options) {
		o.logger = l
	}
}
