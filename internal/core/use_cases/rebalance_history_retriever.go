package use_cases

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/ports"
)

type RebalanceHistoryRetriever interface {
	RetrieveRebalanceHistory(ctx context.Context, address common.Address) ([]domain.RebalanceNote, error)
}

type rebalanceHistoryRetriever struct {
	vaultRepository ports.VaultRepository
}

func NewRebalanceHistoryRetriever(vaultRepository ports.VaultRepository) RebalanceHistoryRetriever {
	return &rebalanceHistoryRetriever{
		vaultRepository: vaultRepository,
	}
}

func (p *rebalanceHistoryRetriever) RetrieveRebalanceHistory(ctx context.Context, address common.Address) ([]domain.RebalanceNote, error) {
	return p.vaultRepository.GetRebalanceHistory(ctx, address)
}
