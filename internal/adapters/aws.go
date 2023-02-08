package adapters

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/protohedge/protohedge.api/internal/config"
)

type AwsClient struct {
	SchedulerClient  *scheduler.Client
	EventsClient     *cloudwatchevents.Client
	CloudwatchClient *cloudwatch.Client
	Config           awsConfig.Config
}

func NewClient(config *config.Config) *AwsClient {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(config.AwsRegion))

	if err != nil {
		panic(err)
	}

	return &AwsClient{
		Config:           cfg,
		SchedulerClient:  scheduler.NewFromConfig(cfg),
		EventsClient:     cloudwatchevents.NewFromConfig(cfg),
		CloudwatchClient: cloudwatch.NewFromConfig(cfg),
	}
}
