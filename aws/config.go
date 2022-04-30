package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
)

func GetConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithDefaultRegion("ap-northeast-2"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
