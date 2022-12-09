package http_server_dtos

import "math/big"

type TokenAllocationResponseDto struct {
	Percentage      *big.Int `json:"percentage"`
	TokenAddress    string   `json:"tokenAddress"`
	Symbol          string   `json:"symbol"`
	Leverage        *big.Int `json:"leverage"`
	CollateralRatio *big.Int `json:"collateralRatio"`
}
