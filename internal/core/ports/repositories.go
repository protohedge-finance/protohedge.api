package ports

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/protohedge/protohedge.api/internal/core/domain"
)

type VaultRepository interface {
	GetVault(ctx context.Context, address common.Address) (*domain.Vault, error)
}
