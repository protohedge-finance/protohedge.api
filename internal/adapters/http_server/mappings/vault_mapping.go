package http_server_mappings

import (
	http_server_dtos "github.com/protohedge/protohedge.api/internal/adapters/http_server/dtos"
	"github.com/protohedge/protohedge.api/internal/core/domain"
)

func ToVaultResponseDto(vault domain.Vault) http_server_dtos.VaultResponseDto {
	positionManagersDtos := []http_server_dtos.PositionManagerResponseDto{}

	for _, p := range vault.PositionManagers {
		positionManagersDtos = append(positionManagersDtos, ToPositionManagerResponseDto(p))
	}

	return http_server_dtos.VaultResponseDto{
		VaultAddress:       vault.VaultAddress.String(),
		PositionManagers:   positionManagersDtos,
		VaultWorth:         vault.VaultWorth.String(),
		AvailableLiquidity: vault.AvailableLiquidity.String(),
		Pnl:                vault.Pnl.String(),
	}
}

func ToPnlTimestampDto(historicPnl domain.TimePoint) http_server_dtos.TimePointDto {
	return http_server_dtos.TimePointDto{
		Timestamp: historicPnl.Timestamp,
		Point:     historicPnl.Point.String(),
	}
}

func ToHistoricPnlDto(historicPnl []domain.TimePoint) *http_server_dtos.HistoricPnlResponseDto {
	historicPnlDto := []http_server_dtos.TimePointDto{}

	for _, p := range historicPnl {
		historicPnlDto = append(historicPnlDto, ToPnlTimestampDto(p))
	}

	return &http_server_dtos.HistoricPnlResponseDto{
		HistoricPnl: historicPnlDto,
	}
}

func ToRebalanceNotesDto(rebalanceNotes []domain.RebalanceNote) *http_server_dtos.RebalanceNotesResponseDto {
	rebalanceNotesResponseDto := []http_server_dtos.RebalanceNoteResponseDto{}

	for _, p := range rebalanceNotes {
		rebalanceNotesResponseDto = append(rebalanceNotesResponseDto, http_server_dtos.RebalanceNoteResponseDto{
			Note: p.Note,
			Date: p.Date,
		})
	}

	return &http_server_dtos.RebalanceNotesResponseDto{
		RebalanceNotes: rebalanceNotesResponseDto,
	}
}

func ToRebalanceInfoDto(rebalanceInfo domain.RebalanceInfo) *http_server_dtos.RebalanceInfoResponseDto {
	return &http_server_dtos.RebalanceInfoResponseDto{
		RebalanceIntervalSeconds: int(rebalanceInfo.RebalanceInterval.Seconds()),
		DurationRemainingSeconds: int(rebalanceInfo.DurationRemaining.Seconds()),
	}
}
