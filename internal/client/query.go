package client

import (
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
	"google.golang.org/grpc"
)

type QueryClient struct {
	Position api.PositionServiceClient
	Price    api.PriceServiceClient
	Glp      api.GlpServiceClient
}

func NewQueryClient() (*QueryClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	positionClient := api.NewPositionServiceClient(conn)
	priceClient := api.NewPriceServiceClient(conn)
	glpClient := api.NewGlpServiceClient(conn)

	queryClient := &QueryClient{
		Position: positionClient,
		Price:    priceClient,
		Glp:      glpClient,
	}

	return queryClient, conn, nil
}
