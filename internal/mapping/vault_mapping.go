package mapping

import (
	"github.com/protohedge/protohedge.api/internal/contracts/vault_contract"
	"github.com/protohedge/protohedge.api/internal/models"
	"github.com/protohedge/protohedge.api/internal/models/dto"
)

func ToVaultModel(vault vault_contract.VaultStats) models.Vault {
	positionManagers := []models.PositionManager{}

	for _, positionManager := range vault.PositionManagers {
		positionManagers = append(positionManagers, ToVaultPositionManagerModel(positionManager))
	}

	return models.Vault{
		VaultAddress:       vault.VaultAddress,
		PositionManagers:   positionManagers,
		VaultWorth:         vault.VaultWorth,
		AvailableLiquidity: vault.AvailableLiquidity,
	}
}

func ToVaultResponseDto(vault *models.Vault) *dto.VaultResponseDto {
	positionManagersDtos := []dto.PositionManagerResponseDto{}

	for _, p := range vault.PositionManagers {
		positionManagersDtos = append(positionManagersDtos, ToPositionManagerResponseDto(p))
	}

	return &dto.VaultResponseDto{
		VaultAddress:       vault.VaultAddress.String(),
		PositionManagers:   positionManagersDtos,
		VaultWorth:         vault.VaultWorth.String(),
		AvailableLiquidity: vault.AvailableLiquidity.String(),
	}
}
