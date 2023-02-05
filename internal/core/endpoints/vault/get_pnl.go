package vault_endpoints

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/endpoint"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
)

type GetPnlRequest struct {
	Address string
}

type GetPnlResponse struct {
	HistoricPnl []domain.TimePoint
}

func NewGetHistoricPnl(pnlRetriever use_cases.PnlRetriever) endpoint.Endpoint {
	var getHistoricPnl endpoint.Endpoint
	{
		getHistoricPnl = CreateGetHistoricPnl(pnlRetriever)

	}

	return getHistoricPnl
}

func CreateGetHistoricPnl(pnlRetriever use_cases.PnlRetriever) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPnlRequest)
		historicPnl, err := pnlRetriever.RetrieveHistoricPnl(ctx, common.HexToAddress(req.Address))

		return GetPnlResponse{
			HistoricPnl: historicPnl,
		}, err
	}
}
