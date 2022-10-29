package dto

type PositionManagerDto struct {
	PnlWorth      string `json:"pnlWorth"`
	PnlPercentage string `json:"pnlPercentage"`
	CostBasis     string `json:"costBasis"`
	PositionWorth string `json:"positionWorth"`
}
