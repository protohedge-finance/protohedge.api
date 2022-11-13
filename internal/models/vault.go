package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Vault struct {
	VaultAddress       common.Address
	PositionManagers   []PositionManager
	VaultWorth         *big.Int
	AvailableLiquidity *big.Int
}
