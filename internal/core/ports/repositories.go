package ports

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/protohedge/protohedge.api/internal/core/domain"
)

type VaultRepository interface {
	GetVault(context.Context, common.Address) (*domain.Vault, error)
	GetHistoricVaultPnl(context.Context, common.Address) ([]domain.TimePoint, error)
	GetRebalanceHistory(context.Context, common.Address) ([]domain.RebalanceNote, error)
}
