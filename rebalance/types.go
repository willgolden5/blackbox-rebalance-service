package rebalance

import (
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	nt, err := time.Parse("2006-01-02T15:04:05.999", s)
	*ct = CustomTime(nt)
	return
}

type ActiveStrategy struct {
	ActivatedAt CustomTime `json:"activatedAt"`
	AlpacaID    string     `json:"alpacaid"`
	Amount      float64    `json:"amount"`
	ID          string     `json:"id"`
	StrategyID  string     `json:"strategyId"`
	UpdatedAt   CustomTime `json:"updatedAt"`
	UserID      string     `json:"userId"`
}

type Position struct {
	Ticker string          `json:"ticker"`
	NAV    decimal.Decimal `json:"navPercentage"`
}

type AlpacaPositions struct {
	AlpacaID  string `json:"alpacaId"`
	AccountID string `json:"accountId"`
	Positions []Position
}

type PortDiff struct {
	Ticker       string
	PercentOfNav decimal.Decimal
	Difference   interface{}
}
