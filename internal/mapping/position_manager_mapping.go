package mapping

import (
	"github.com/protohedge/protohedge.api/internal/contracts/position_manager_contract"
	"github.com/protohedge/protohedge.api/internal/contracts/vault_contract"
	"github.com/protohedge/protohedge.api/internal/models"
	"github.com/protohedge/protohedge.api/internal/models/dto"
)

func ToPositionManagerModel(positionManager position_manager_contract.PositionManagerStats) models.PositionManager {
	tokenExposures := []*models.TokenExposure{}
	tokenAllocations := []*models.TokenAllocation{}

	for _, exposure := range positionManager.TokenExposures {
		tokenExposures = append(tokenExposures, &models.TokenExposure{
			Symbol: exposure.Symbol,
			Amount: exposure.Amount,
			Token:  exposure.Token.String(),
		})
	}

	for _, allocation := range positionManager.TokenAllocation {
		tokenAllocations = append(tokenAllocations, &models.TokenAllocation{
			Symbol:       allocation.Symbol,
			Percentage:   allocation.Percentage,
			TokenAddress: allocation.TokenAddress,
			Leverage:     allocation.Leverage,
		})
	}

	return models.PositionManager{
		PositionManagerAddress: positionManager.PositionManagerAddress,
		PositionWorth:          positionManager.PositionWorth,
		CostBasis:              positionManager.CostBasis,
		Pnl:                    positionManager.Pnl,
		TokenExposures:         tokenExposures,
		TokenAllocation:        tokenAllocations,
		CanRebalance:           positionManager.CanRebalance,
	}
}

func ToVaultPositionManagerModel(positionManager vault_contract.PositionManagerStats) models.PositionManager {
	tokenExposures := []*models.TokenExposure{}
	tokenAllocations := []*models.TokenAllocation{}

	for _, exposure := range positionManager.TokenExposures {
		tokenExposures = append(tokenExposures, &models.TokenExposure{
			Symbol: exposure.Symbol,
			Amount: exposure.Amount,
			Token:  exposure.Token.String(),
		})
	}

	for _, allocation := range positionManager.TokenAllocation {
		tokenAllocations = append(tokenAllocations, &models.TokenAllocation{
			Symbol:       allocation.Symbol,
			Percentage:   allocation.Percentage,
			TokenAddress: allocation.TokenAddress,
			Leverage:     allocation.Leverage,
		})
	}

	return models.PositionManager{
		PositionManagerAddress: positionManager.PositionManagerAddress,
		PositionWorth:          positionManager.PositionWorth,
		CostBasis:              positionManager.CostBasis,
		Pnl:                    positionManager.Pnl,
		TokenExposures:         tokenExposures,
		TokenAllocation:        tokenAllocations,
		CanRebalance:           positionManager.CanRebalance,
	}
}

func ToPositionManagerModels(positionManagersStats []position_manager_contract.PositionManagerStats) []models.PositionManager {
	positionManagers := []models.PositionManager{}

	for i, positionManager := range positionManagersStats {
		positionManagers[i] = ToPositionManagerModel(positionManager)
	}

	return positionManagers
}

func ToPositionManagerResponseDto(positionManager models.PositionManager) dto.PositionManagerResponseDto {
	tokenExposures := []dto.TokenExposureResponseDto{}
	tokenAllocations := []dto.TokenAllocationResponseDto{}

	for _, exposure := range positionManager.TokenExposures {
		tokenExposures = append(tokenExposures, dto.TokenExposureResponseDto{
			Symbol: exposure.Symbol,
			Amount: exposure.Amount.String(),
			Token:  exposure.Token,
		})
	}

	for _, allocation := range positionManager.TokenAllocation {
		tokenAllocations = append(tokenAllocations, dto.TokenAllocationResponseDto{
			Symbol:       allocation.Symbol,
			Percentage:   allocation.Percentage,
			TokenAddress: allocation.TokenAddress.String(),
			Leverage:     allocation.Leverage,
		})
	}

	return dto.PositionManagerResponseDto{
		PositionManagerAddress: positionManager.PositionManagerAddress.String(),
		PositionWorth:          positionManager.PositionWorth.String(),
		CostBasis:              positionManager.CostBasis.String(),
		Pnl:                    positionManager.Pnl.String(),
		TokenExposures:         tokenExposures,
		TokenAllocation:        tokenAllocations,
		CanRebalance:           positionManager.CanRebalance,
	}
}
