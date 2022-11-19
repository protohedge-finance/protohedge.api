package contract_repositories

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/abi/vault_contract"
	contract_mappings "github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/mappings"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/ports"
)

type vaultRepository struct {
	ethClient *ethclient.Client
}

func NewVaultRepository(ethClient *ethclient.Client) ports.VaultRepository {
	return &vaultRepository{
		ethClient: ethClient,
	}
}

func (v *vaultRepository) GetVault(ctx context.Context, address common.Address) (*domain.Vault, error) {
	ctr, err := vault_contract.NewVaultContract(address, v.ethClient)

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	result, err := ctr.Stats(callOpts)

	if err != nil {
		return nil, err
	}

	vault := contract_mappings.ToVaultModel(result)
	return &vault, nil
}
