package domain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenExposure struct {
	Symbol string   `json:"symbol"`
	Amount *big.Int `json:"amount"`
	Token  string   `json:"token"`
}

type TokenAllocation struct {
	Symbol          string         `json:"symbol"`
	Percentage      *big.Int       `json:"percentage"`
	TokenAddress    common.Address `json:"tokenAddress"`
	Leverage        *big.Int       `json:"leverage"`
	CollateralRatio *big.Int       `json:"collateralRatio"`
}

type PositionManager struct {
	PositionManagerAddress common.Address     `json:"positionManagerAddress"`
	Name                   string             `json:"name"`
	PositionWorth          *big.Int           `json:"positionWorth"`
	CostBasis              *big.Int           `json:"costBasis"`
	Pnl                    *big.Int           `json:"pnl"`
	CollateralRatio        *big.Int           `json:"collateralRatio"`
	LoanWorth              *big.Int           `json:"loanWorth"`
	LiquidationLevel       *big.Int           `json:"liquidationLevel"`
	Collateral             *big.Int           `json:"collateral"`
	TokenExposures         []*TokenExposure   `json:"tokenExposures"`
	TokenAllocation        []*TokenAllocation `json:"tokenAllocation"`
	CanRebalance           bool               `json:"canRebalance"`
}
