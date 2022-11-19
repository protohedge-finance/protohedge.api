package use_cases

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/ports"
)

type VaultRetriever interface {
	RetrieveVault(ctx context.Context, address common.Address) (*domain.Vault, error)
}

type vaultRetriever struct {
	vaultRepository ports.VaultRepository
}

func NewVaultRetriever(vaultRepository ports.VaultRepository) VaultRetriever {
	return &vaultRetriever{
		vaultRepository: vaultRepository,
	}
}

func (s *vaultRetriever) RetrieveVault(ctx context.Context, address common.Address) (*domain.Vault, error) {
	return s.vaultRepository.GetVault(ctx, address)
}
