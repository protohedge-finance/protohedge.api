package domain

import "time"

type RebalanceInfo struct {
	RebalanceInterval time.Duration
	DurationRemaining time.Duration
}
