package contract_mappings

import (
	"github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/abi/position_manager_contract"
	"github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/abi/vault_contract"
	"github.com/protohedge/protohedge.api/internal/core/domain"
)

func ToPositionManagerModel(positionManager position_manager_contract.PositionManagerStats) domain.PositionManager {
	tokenExposures := []*domain.TokenExposure{}
	tokenAllocations := []*domain.TokenAllocation{}

	for _, exposure := range positionManager.TokenExposures {
		tokenExposures = append(tokenExposures, &domain.TokenExposure{
			Symbol: exposure.Symbol,
			Amount: exposure.Amount,
			Token:  exposure.Token.String(),
		})
	}

	for _, allocation := range positionManager.TokenAllocations {
		tokenAllocations = append(tokenAllocations, &domain.TokenAllocation{
			Symbol:       allocation.Symbol,
			Percentage:   allocation.Percentage,
			TokenAddress: allocation.TokenAddress,
			Leverage:     allocation.Leverage,
		})
	}
	return domain.PositionManager{
		PositionManagerAddress: positionManager.PositionManagerAddress,
		Name:                   positionManager.Name,
		PositionWorth:          positionManager.PositionWorth,
		CostBasis:              positionManager.CostBasis,
		Pnl:                    positionManager.Pnl,
		CollateralRatio:        positionManager.CollateralRatio,
		LoanWorth:              positionManager.LoanWorth,
		Collateral:             positionManager.Collateral,
		TokenExposures:         tokenExposures,
		TokenAllocation:        tokenAllocations,
	}
}

func ToVaultPositionManagerModel(positionManager vault_contract.PositionManagerStats) domain.PositionManager {
	tokenExposures := []*domain.TokenExposure{}
	tokenAllocations := []*domain.TokenAllocation{}

	for _, exposure := range positionManager.TokenExposures {
		tokenExposures = append(tokenExposures, &domain.TokenExposure{
			Symbol: exposure.Symbol,
			Amount: exposure.Amount,
			Token:  exposure.Token.String(),
		})
	}

	for _, allocation := range positionManager.TokenAllocations {
		tokenAllocations = append(tokenAllocations, &domain.TokenAllocation{
			Symbol:       allocation.Symbol,
			Percentage:   allocation.Percentage,
			TokenAddress: allocation.TokenAddress,
			Leverage:     allocation.Leverage,
		})
	}

	return domain.PositionManager{
		PositionManagerAddress: positionManager.PositionManagerAddress,
		Name:                   positionManager.Name,
		PositionWorth:          positionManager.PositionWorth,
		CostBasis:              positionManager.CostBasis,
		Pnl:                    positionManager.Pnl,
		CollateralRatio:        positionManager.CollateralRatio,
		LoanWorth:              positionManager.LoanWorth,
		LiquidationLevel:       positionManager.LiquidationLevel,
		Collateral:             positionManager.Collateral,
		TokenExposures:         tokenExposures,
		TokenAllocation:        tokenAllocations,
	}
}

func ToPositionManagerModels(positionManagersStats []position_manager_contract.PositionManagerStats) []domain.PositionManager {
	positionManagers := []domain.PositionManager{}

	for i, positionManager := range positionManagersStats {
		positionManagers[i] = ToPositionManagerModel(positionManager)
	}

	return positionManagers
}
