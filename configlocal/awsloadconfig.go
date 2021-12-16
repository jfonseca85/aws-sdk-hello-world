package configlocal

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/viper"
)

func awsloadconfig(ctx context.Context, viper *viper.Viper) (aws.Config, error) {

	awsEndpoint := viper.GetString("aws.endpoint_url")
	awsRegion := viper.GetString("aws.default_region")
	//TODO: Inserir validação caso não exista
	//os.Setenv("AWS_SECRET_ACCESS_KEY", viper.GetString("aws.secret_access_key"))
	//os.Setenv("AWS_ACCESS_KEY_ID", viper.GetString("aws.access_key_id"))

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
