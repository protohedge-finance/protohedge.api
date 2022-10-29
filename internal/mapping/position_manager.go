package mapping

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/protohedge/protohedge.api/internal/contracts/position_manager_contract"
	"github.com/protohedge/protohedge.api/internal/models"
)

func ToPositionManagerModel(positionManager position_manager_contract.PositionStats, address common.Address) models.PositionManager {
	tokenExposures := []*models.TokenExposure{}
	tokenAllocations := []*models.TokenAllocation{}

	for _, exposure := range positionManager.TokenExposures {
		tokenExposures = append(tokenExposures, &models.TokenExposure{
			Amount: exposure.Amount,
			Token:  exposure.Token.String(),
		})
	}

	for _, allocation := range positionManager.TokenAllocation {
		tokenAllocations = append(tokenAllocations, &models.TokenAllocation{
			Percentage:   allocation.Percentage,
			TokenAddress: allocation.TokenAddress,
			Leverage:     allocation.Leverage,
		})
	}

	return models.PositionManager{
		Address:         address,
		PositionWorth:   positionManager.PositionWorth,
		CostBasis:       positionManager.CostBasis,
		Pnl:             positionManager.Pnl,
		TokenExposures:  tokenExposures,
		TokenAllocation: tokenAllocations,
		CanRebalance:    positionManager.CanRebalance,
	}
}
