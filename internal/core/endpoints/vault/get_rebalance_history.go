package vault_endpoints

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/endpoint"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
)

type GetRebalanceNotesRequest struct {
	Address string
}

type GetRebalanceNotesResponse struct {
	RebalanceNotes []domain.RebalanceNote
}

func NewGetRebalanceNotes(rebalanceNotesRetriever use_cases.RebalanceNotesRetriever) endpoint.Endpoint {
	var getRebalanceNotes endpoint.Endpoint
	{
		getRebalanceNotes = CreateGetRebalanceNotes(rebalanceNotesRetriever)

	}

	return getRebalanceNotes
}

func CreateGetRebalanceNotes(rebalanceNotesRetriever use_cases.RebalanceNotesRetriever) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRebalanceNotesRequest)
		rebalanceNotes, err := rebalanceNotesRetriever.RetrieveRebalanceNotes(ctx, common.HexToAddress(req.Address))

		return GetRebalanceNotesResponse{
			RebalanceNotes: rebalanceNotes,
		}, err
	}
}
