package configlocal

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/viper"
)

type configlocal struct {
	Viper     *viper.Viper
	AWSClient awsclient
}

type awsclient struct {
	config aws.Config
}

func newAWSClient(config aws.Config) awsclient {
	return awsclient{
		config: config,
	}
}

func NewConfig(ctx context.Context) (*configlocal, error) {
	viper.AddConfigPath(".")  // to work on dev and production envs
	viper.AddConfigPath("./") // to work on dev and production envs
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	awsconfig, err := aWSConfig(ctx, viper.GetViper())

	if err != nil {
		return nil, err
	}

	awsclient := newAWSClient(awsconfig)

	return &configlocal{
		Viper:     viper.GetViper(),
		AWSClient: awsclient,
	}, nil
}

func aWSConfig(ctx context.Context, viper *viper.Viper) (aws.Config, error) {

	awsEndpoint := viper.GetString("aws.endpoint_url")
	awsRegion := viper.GetString("aws.default_region")

	os.Setenv("AWS_SECRET_ACCESS_KEY", viper.GetString("aws.secret_access_key"))
	os.Setenv("AWS_ACCESS_KEY_ID", viper.GetString("aws.access_key_id"))

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" && awsRegion != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	instance, err := config.LoadDefaultConfig(ctx, config.WithEndpointResolver(customResolver))

	return instance, err
}
