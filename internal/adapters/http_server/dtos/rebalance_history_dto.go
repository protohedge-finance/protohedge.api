package http_server_dtos

import (
	"time"
)

type RebalanceHistoryResponseDto struct {
	RebalanceHistory []RebalanceNoteResponseDto `json:"rebalanceHistory"`
}

type RebalanceNoteResponseDto struct {
	Note string    `json:"note"`
	Date time.Time `json:"date"`
}
