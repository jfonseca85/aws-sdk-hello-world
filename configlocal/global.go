package configlocal

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func InitAWSConfig(cfg *configlocal) (aws.Config, error) {

	awsEndpoint := cfg.Viper.GetString("aws.endpoint_url")
	awsRegion := cfg.Viper.GetString("aws.default_region")

	os.Setenv("AWS_SECRET_ACCESS_KEY", cfg.Viper.GetString("aws.secret_access_key"))
	os.Setenv("AWS_ACCESS_KEY_ID", cfg.Viper.GetString("aws.access_key_id"))

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

	instance, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(customResolver))

	return instance, err
}
