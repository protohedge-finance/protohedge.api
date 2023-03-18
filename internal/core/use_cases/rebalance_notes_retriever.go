package use_cases

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/ports"
)

type RebalanceNotesRetriever interface {
	RetrieveRebalanceNotes(ctx context.Context, address common.Address) ([]domain.RebalanceNote, error)
}

type rebalanceNotesRetriever struct {
	vaultRepository ports.VaultRepository
}

func NewRebalanceNotesRetriever(vaultRepository ports.VaultRepository) RebalanceNotesRetriever {
	return &rebalanceNotesRetriever{
		vaultRepository: vaultRepository,
	}
}

func (p *rebalanceNotesRetriever) RetrieveRebalanceNotes(ctx context.Context, address common.Address) ([]domain.RebalanceNote, error) {
	return p.vaultRepository.GetRebalanceNotes(ctx, address)
}
