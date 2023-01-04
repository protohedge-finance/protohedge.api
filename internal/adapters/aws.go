package adapters

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

type AwsClient struct {
	SchedulerClient *scheduler.Client
	Config          config.Config
}

func NewClient() *AwsClient {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic(err)
	}

	return &AwsClient{
		Config:          cfg,
		SchedulerClient: scheduler.NewFromConfig(cfg),
	}
}
