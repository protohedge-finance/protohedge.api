package http_server_dtos

type RebalanceInfoResponseDto struct {
	RebalanceIntervalSeconds int `json:"rebalanceIntervalSeconds"`
	DurationRemainingSeconds int `json:"durationRemainingSeconds"`
}
