package http_server_dtos

type PositionManagerResponseDto struct {
	PositionManagerAddress string                       `json:"positionManagerAddress"`
	Name                   string                       `json:"name"`
	PositionWorth          string                       `json:"positionWorth"`
	CostBasis              string                       `json:"costBasis"`
	Pnl                    string                       `json:"pnl"`
	CollateralRatio        string                       `json:"collateralRatio"`
	LoanWorth              string                       `json:"loanWorth"`
	LiquidationLevel       string                       `json:"liquidationLevel"`
	Collateral             string                       `json:"collateral"`
	TokenExposures         []TokenExposureResponseDto   `json:"tokenExposures"`
	TokenAllocation        []TokenAllocationResponseDto `json:"tokenAllocation"`
	CanRebalance           bool                         `json:"canRebalance"`
}
