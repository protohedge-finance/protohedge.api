package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenExposure struct {
	Amount *big.Int
	Token  string
}

type TokenAllocation struct {
	Percentage   *big.Int
	TokenAddress common.Address
	Leverage     *big.Int
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
