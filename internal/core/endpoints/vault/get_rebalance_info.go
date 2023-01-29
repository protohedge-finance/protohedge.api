package vault_endpoints

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/endpoint"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
)

type GetRebalanceInfoRequest struct {
	Address string
}

type GetRebalanceInfoResponse struct {
	RebalanceInfo domain.RebalanceInfo
	Err           error
}

func NewGetRebalanceInfoEndpoint(rebalanceInfoRetriever use_cases.RebalanceInfoRetriever) endpoint.Endpoint {
	var getRebalanceInfoEndpoint endpoint.Endpoint
	{
		getRebalanceInfoEndpoint = CreateGetRebalanceInfo(rebalanceInfoRetriever)

	}

	return getRebalanceInfoEndpoint
}

func CreateGetRebalanceInfo(rebalanceInfoRetriever use_cases.RebalanceInfoRetriever) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRebalanceInfoRequest)
		rebalanceInfo, err := rebalanceInfoRetriever.RetrieveRebalanceInfo(ctx, common.HexToAddress(req.Address))
		return GetRebalanceInfoResponse{
			RebalanceInfo: rebalanceInfo,
			Err:           err,
		}, err
	}
}
