package http_server_mappings

import (
	http_server_dtos "github.com/protohedge/protohedge.api/internal/adapters/http_server/dtos"
	"github.com/protohedge/protohedge.api/internal/core/domain"
)

func ToPositionManagerResponseDto(positionManager domain.PositionManager) http_server_dtos.PositionManagerResponseDto {
	tokenExposures := []http_server_dtos.TokenExposureResponseDto{}
	tokenAllocations := []http_server_dtos.TokenAllocationResponseDto{}

	for _, exposure := range positionManager.TokenExposures {
		tokenExposures = append(tokenExposures, http_server_dtos.TokenExposureResponseDto{
			Symbol: exposure.Symbol,
			Amount: exposure.Amount.String(),
			Token:  exposure.Token,
		})
	}

	for _, allocation := range positionManager.TokenAllocation {
		tokenAllocations = append(tokenAllocations, http_server_dtos.TokenAllocationResponseDto{
			Symbol:       allocation.Symbol,
			Percentage:   allocation.Percentage,
			TokenAddress: allocation.TokenAddress.String(),
			Leverage:     allocation.Leverage,
		})
	}

	return http_server_dtos.PositionManagerResponseDto{
		PositionManagerAddress: positionManager.PositionManagerAddress.String(),
		PositionWorth:          positionManager.PositionWorth.String(),
		CostBasis:              positionManager.CostBasis.String(),
		Pnl:                    positionManager.Pnl.String(),
		TokenExposures:         tokenExposures,
		TokenAllocation:        tokenAllocations,
		CanRebalance:           positionManager.CanRebalance,
	}
}
