package domain

import (
	"math/big"
	"time"
)

type TimePoint struct {
	Timestamp time.Time
	Point     *big.Int
}
