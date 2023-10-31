package config

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func ConfigStorage() *s3.Client {

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
				URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", viper.GetString("storage.ACCOUNT_ID")),
			},
			nil
	})

	configuration, ErrConfiguration := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(viper.GetString("storage.ACCESS_KEY_ID"), viper.GetString("storage.SECRET_ACCESS_KEY"), "")),
	)

	if ErrConfiguration != nil {
		panic(ErrConfiguration.Error())
	}

	client := s3.NewFromConfig(configuration)

	return client
}
