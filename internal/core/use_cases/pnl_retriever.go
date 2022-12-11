package use_cases

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/ports"
)

type PnlRetriever interface {
	RetrieveHistoricPnl(ctx context.Context, address common.Address) ([]domain.TimePoint, error)
}

type pnlRetriever struct {
	vaultRepository ports.VaultRepository
}

func NewPnlRetriever(vaultRepository ports.VaultRepository) PnlRetriever {
	return &pnlRetriever{
		vaultRepository: vaultRepository,
	}
}

func (p *pnlRetriever) RetrieveHistoricPnl(ctx context.Context, address common.Address) ([]domain.TimePoint, error) {
	return p.vaultRepository.GetHistoricVaultPnl(ctx, address)
}
