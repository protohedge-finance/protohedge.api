package dto

type PositionManagerResponseDto struct {
	Address         string                       `json:"address"`
	PositionWorth   string                       `json:"positionWorth"`
	CostBasis       string                       `json:"costBasis"`
	Pnl             string                       `json:"pnl"`
	TokenExposures  []TokenExposureResponseDto   `json:"tokenExposures"`
	TokenAllocation []TokenAllocationResponseDto `json:"tokenAllocation"`
	CanRebalance    bool                         `json:"canRebalance"`
}
