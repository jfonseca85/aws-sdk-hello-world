package configlocal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/viper"
	"strings"
)

type viperloadconfig struct {
	Viper     *viper.Viper
	AWSClient aws.Config
}

func NewConfig(ctx context.Context) (*viperloadconfig, error) {
	viper.AddConfigPath(".")  // to work on dev and production envs
	viper.AddConfigPath("./") // to work on dev and production envs
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	awsconfig, err := awsloadconfig(ctx, viper.GetViper())
	if err != nil {
		return nil, err
	}

	return &viperloadconfig{
		Viper:     viper.GetViper(),
		AWSClient: awsconfig,
	}, nil
}
