package vault

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/protohedge/protohedge.api/internal/contracts/vault_contract"
	"github.com/protohedge/protohedge.api/internal/mapping"
	"github.com/protohedge/protohedge.api/internal/models"
)

type Service interface {
	GetVault(address common.Address) (*models.Vault, error)
}

type vaultService struct {
	ethClient *ethclient.Client
}

func NewVaultService(ethClient *ethclient.Client) Service {
	return &vaultService{
		ethClient: ethClient,
	}
}

func (s *vaultService) GetVault(address common.Address) (*models.Vault, error) {
	ctr, err := vault_contract.NewVaultContract(address, s.ethClient)

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	result, err := ctr.Stats(callOpts)

	if err != nil {
		return nil, err
	}

	vault := mapping.ToVaultModel(result)
	return &vault, nil
}
