package http_server_dtos

type PositionManagerResponseDto struct {
	PositionManagerAddress string                       `json:"positionManagerAddress"`
	PositionWorth          string                       `json:"positionWorth"`
	CostBasis              string                       `json:"costBasis"`
	Pnl                    string                       `json:"pnl"`
	TokenExposures         []TokenExposureResponseDto   `json:"tokenExposures"`
	TokenAllocation        []TokenAllocationResponseDto `json:"tokenAllocation"`
	CanRebalance           bool                         `json:"canRebalance"`
}
