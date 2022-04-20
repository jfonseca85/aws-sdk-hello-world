package config

import (
	"github.com/jfonseca85/aws-sdk-hello-world/log"
	"github.com/spf13/viper"
)

type Option interface {
	apply(v *viper.Viper)
}

type optionFunc func(v *viper.Viper)

func (fn optionFunc) apply(v *viper.Viper) {
	fn(v)
}

type options struct {
	logger log.Logger
}

// WithSource with config source.
func WithSource(s string) Option {
	return optionFunc(func(v *viper.Viper) {
		v.SetConfigFile(s)
	})
}

// WithPath with config source.
func WithPath(s ...string) Option {
	return optionFunc(func(v *viper.Viper) {
		for _, src := range s {
			v.AddConfigPath(src)
		}
	})
}

// WithPath with config source.
func WithName(s string) Option {
	return optionFunc(func(v *viper.Viper) {
		v.SetConfigName(s)
	})
}

// WithPath with config source.
func WithType(s string) Option {
	return optionFunc(func(v *viper.Viper) {
		v.SetConfigType(s)
	})
}
