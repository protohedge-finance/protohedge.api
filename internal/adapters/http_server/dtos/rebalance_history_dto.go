package http_server_dtos

import (
	"time"
)

type RebalanceNotesResponseDto struct {
	RebalanceNotes []RebalanceNoteResponseDto `json:"rebalanceNotes"`
}

type RebalanceNoteResponseDto struct {
	Note string    `json:"note"`
	Date time.Time `json:"date"`
}
