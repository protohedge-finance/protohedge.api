package vault_endpoints

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/endpoint"
	http_server_dtos "github.com/protohedge/protohedge.api/internal/adapters/http_server/dtos"
	http_server_mappings "github.com/protohedge/protohedge.api/internal/adapters/http_server/mappings"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
)

type GetVaultRequest struct {
	Address string
}

type GetVaultResponse struct {
	Vault domain.Vault
	Err   error
}

type GetHistoricVaultPnlRequest struct {
	Address string
}

type GetHistoricVaultPnlResponse struct {
	HistoricVaultPnl []http_server_dtos.TimePointDto
	Err              error
}

func NewGetVaultEndpoint(vaultRetriever use_cases.VaultRetriever) endpoint.Endpoint {
	var getVaultEndpoint endpoint.Endpoint
	{
		getVaultEndpoint = CreateGetVaultEndpoint(vaultRetriever)

	}

	return getVaultEndpoint
}

func CreateGetVaultEndpoint(vaultRetriever use_cases.VaultRetriever) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetVaultRequest)
		vault, err := vaultRetriever.RetrieveVault(ctx, common.HexToAddress(req.Address))

		return GetVaultResponse{
			Vault: *vault,
			Err:   err,
		}, nil
	}
}

func CreateGetHistoricVaultPnl(pnlRetriever use_cases.PnlRetriever) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetHistoricVaultPnlRequest)

		historicPnl, err := pnlRetriever.RetrieveHistoricPnl(ctx, common.HexToAddress(req.Address))

		if err != nil {
			return GetHistoricVaultPnlResponse{
				Err: err,
			}, nil
		}

		historicPnlDto := http_server_mappings.ToHistoricPnlDto(historicPnl)

		return GetHistoricVaultPnlResponse{
			HistoricVaultPnl: historicPnlDto.HistoricPnl,
		}, nil
	}
}
