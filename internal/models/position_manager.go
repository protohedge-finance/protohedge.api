package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenExposure struct {
	Amount *big.Int `json:"amount"`
	Token  string   `json:"token"`
}

type TokenAllocation struct {
	Percentage   *big.Int       `json:"percentage"`
	TokenAddress common.Address `json:"tokenAddress"`
	Leverage     *big.Int       `json:"leverage"`
}

type PositionManager struct {
	Address         common.Address     `json:"address"`
	PositionWorth   *big.Int           `json:"positionWorth"`
	CostBasis       *big.Int           `json:"costBasis"`
	Pnl             *big.Int           `json:"pnl"`
	TokenExposures  []*TokenExposure   `json:"tokenExposures"`
	TokenAllocation []*TokenAllocation `json:"tokenAllocation"`
	CanRebalance    bool               `json:"canRebalance"`
}
