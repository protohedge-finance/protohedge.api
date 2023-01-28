package vault_endpoints

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/endpoint"
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
			Vault: vault,
			Err:   err,
		}, nil
	}
}
