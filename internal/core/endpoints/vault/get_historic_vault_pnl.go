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
}

func NewGetHistoricVaultPnl(pnlRetriever use_cases.PnlRetriever) endpoint.Endpoint {
	var getRebalanceNotes endpoint.Endpoint
	{
		getRebalanceNotes = CreateGetHistoricVaultPnl(pnlRetriever)
	}

	return getRebalanceNotes
}

func CreateGetHistoricVaultPnl(pnlRetriever use_cases.PnlRetriever) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetHistoricVaultPnlRequest)
		historicPnl, err := pnlRetriever.RetrieveHistoricPnl(ctx, common.HexToAddress(req.Address))
		return GetHistoricVaultPnlResponse{
			HistoricVaultPnl: historicPnl,
		}, err
	}
}
