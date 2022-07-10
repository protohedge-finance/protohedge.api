package models

import "math/big"

type Pnl struct {
	Amount     *big.Int
	Percentage *big.Float
}
