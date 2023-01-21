package vault_endpoints

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/endpoint"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
)

type GetRebalanceHistoryRequest struct {
	Address string
}

type GetRebalanceHistoryResponse struct {
	RebalanceHistory []domain.RebalanceNote
	Err              error
}

func NewGetRebalanceHistory(rebalanceHistoryRetriever use_cases.RebalanceHistoryRetriever) endpoint.Endpoint {
	var getRebalanceHistory endpoint.Endpoint
	{
		getRebalanceHistory = CreateGetRebalanceHistory(rebalanceHistoryRetriever)

	}

	return getRebalanceHistory
}

func CreateGetRebalanceHistory(rebalanceHistoryRetriever use_cases.RebalanceHistoryRetriever) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRebalanceHistoryRequest)
		rebalanceHistory, err := rebalanceHistoryRetriever.RetrieveRebalanceHistory(ctx, common.HexToAddress(req.Address))

		return GetRebalanceHistoryResponse{
			RebalanceHistory: rebalanceHistory,
			Err:              err,
		}, nil
	}
}
