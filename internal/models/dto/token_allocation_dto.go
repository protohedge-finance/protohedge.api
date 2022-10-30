package dto

import "math/big"

type TokenAllocationResponseDto struct {
	Percentage   *big.Int `json:"percentage"`
	TokenAddress string   `json:"tokenAddress"`
	Leverage     *big.Int `json:"leverage"`
}
