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
	Symbol       string         `json:"symbol"`
	Percentage   *big.Int       `json:"percentage"`
	TokenAddress common.Address `json:"tokenAddress"`
	Leverage     *big.Int       `json:"leverage"`
}

type PositionManager struct {
	PositionManagerAddress common.Address     `json:"positionManagerAddress"`
	PositionWorth          *big.Int           `json:"positionWorth"`
	CostBasis              *big.Int           `json:"costBasis"`
	Pnl                    *big.Int           `json:"pnl"`
	TokenExposures         []*TokenExposure   `json:"tokenExposures"`
	TokenAllocation        []*TokenAllocation `json:"tokenAllocation"`
	CanRebalance           bool               `json:"canRebalance"`
}
