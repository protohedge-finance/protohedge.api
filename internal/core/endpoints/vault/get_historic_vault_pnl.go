package vault_endpoints

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/endpoint"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
)

type GetHistoricVaultPnlRequest struct {
	Address string
}

type GetHistoricVaultPnlResponse struct {
	HistoricVaultPnl []domain.TimePoint
	Err              error
}

func NewGetHistoricVaultPnl(pnlRetriever use_cases.PnlRetriever) endpoint.Endpoint {
	var getRebalanceHistory endpoint.Endpoint
	{
		getRebalanceHistory = CreateGetHistoricVaultPnl(pnlRetriever)
	}

	return getRebalanceHistory
}

func CreateGetHistoricVaultPnl(pnlRetriever use_cases.PnlRetriever) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRebalanceHistoryRequest)
		historicPnl, err := pnlRetriever.RetrieveHistoricPnl(ctx, common.HexToAddress(req.Address))

		return GetHistoricVaultPnlResponse{
			HistoricVaultPnl: historicPnl,
			Err:              err,
		}, nil
	}
}
