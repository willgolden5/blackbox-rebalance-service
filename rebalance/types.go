package rebalance

import (
	"time"
)

type ActiveStrategy struct {
	ActivatedAt time.Time `json:"activatedAt"`
	AlpacaID    string    `json:"alpacaid"`
	Amount      float64   `json:"amount"` // Assuming amount can be a decimal, use float64. Use int if it's always an integer.
	ID          string    `json:"id"`
	StrategyID  string    `json:"strategyId"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserID      string    `json:"userId"`
}
