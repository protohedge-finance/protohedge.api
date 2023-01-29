package use_cases

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/ports"
)

type RebalanceInfoRetriever interface {
	RetrieveRebalanceInfo(ctx context.Context, address common.Address) (domain.RebalanceInfo, error)
}

type rebalanceInfoRetriever struct {
	vaultRepository ports.VaultRepository
}

func NewRebalanceInfoRetriever(vaultRepository ports.VaultRepository) RebalanceInfoRetriever {
	return &rebalanceInfoRetriever{
		vaultRepository: vaultRepository,
	}
}

func (s *rebalanceInfoRetriever) RetrieveRebalanceInfo(ctx context.Context, address common.Address) (domain.RebalanceInfo, error) {
	return s.vaultRepository.GetRebalanceInfo(ctx, address)
}
