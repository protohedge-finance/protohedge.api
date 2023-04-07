package contract_mappings

import (
	"github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/abi/vault_contract"
	"github.com/protohedge/protohedge.api/internal/core/domain"
)

func ToVaultModel(vault vault_contract.VaultStats) domain.Vault {
	positionManagers := []domain.PositionManager{}

	for _, positionManager := range vault.PositionManagers {
		positionManagers = append(positionManagers, ToVaultPositionManagerModel(positionManager))
	}

	return domain.Vault{
		VaultAddress:       vault.VaultAddress,
		PositionManagers:   positionManagers,
		VaultWorth:         vault.VaultWorth,
		AvailableLiquidity: vault.AvailableLiquidity,
		Pnl:                vault.Pnl,
	}
}
