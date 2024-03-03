package rebalance

import (
	"fmt"
	"strings"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/shopspring/decimal"
)

// FIRST:
// get users from ActiveStrategies table where activeStrategyId = activeStrategyId =>
// get users alpaca client with oauth =>

// get the active positions from their alpaca account =>
// create object with alpacaId, accountId, and positions object with nav calculated for each position
// relative to their tradeable balance in their alpaca account

// THEN:
// get the holdings of the strategy specified from the strategy table =>
// compare the strategy holdings to the users current portfolio holdings and output a portDiff struct =>
// if (portDiff is not empty) => sell all equities that need to be sold THEN
// buy all equities that need to be bought =>
// update the user's portfolio in the ActiveStrategies table with rebalance date

func RebalanceUserPortfolio(activeStrategyId string) {
	activeStrategyData, err := getActiveStrategyListing(activeStrategyId)
	if err != nil {
		fmt.Println("error getting users by active strategy", err)
		return
	}
	if activeStrategyData == nil {
		fmt.Println("Active strategy does not exist.", activeStrategyId)
		return
	}
	oAuthId := strings.TrimPrefix(activeStrategyData.AlpacaID, "Bearer ")
	client := alpaca.NewClient(alpaca.ClientOpts{
		OAuth: oAuthId,
	})

	awaitMarketOpen(client)

	account, err := client.GetAccount()
	if err != nil {
		fmt.Println("error getting account", err)
		return
	}

	positions, err := client.GetPositions()
	if err != nil {
		fmt.Println("error getting positions", err)
		return
	}

	currentUserNavs := calculateNAV(positions, account.BuyingPower)
	strategyHoldings, err := getStrategyHoldings(convertToCamelCase(activeStrategyData.StrategyID))

	if err != nil {
		fmt.Println("error getting strategy holdings", err)
		return
	}

	portDiffs := compareHoldings(currentUserNavs, strategyHoldings)
	if len(portDiffs) < 1 {
		fmt.Println("congress buys portfolio is balanced")
	} else {
		// Iterate over portDiffs
		for _, diff := range portDiffs {
			ticker := diff.Ticker
			percentOfNav := diff.PercentOfNav
			difference, ok := diff.Difference.(decimal.Decimal)

			var side string
			var notional decimal.Decimal

			if !ok {
				// Place a trade to buy the percentOfNav of the ticker
				notional = percentOfNav.Mul(account.Cash)
				side = "buy"
			} else {
				// Either buy or sell the difference amount
				notional = difference.Abs().Mul(account.Cash)
				if difference.IsPositive() {
					side = "buy"
				} else {
					side = "sell"
				}
			}

			orderReq := alpaca.PlaceOrderRequest{
				Symbol:      ticker,
				Qty:         &notional,
				Side:        alpaca.Side(side),
				Type:        alpaca.Market,
				TimeInForce: alpaca.Day,
			}

			_, err := client.PlaceOrder(orderReq)
			if err != nil {
				fmt.Println("error placing order", err)
				return
			}
		}

		fmt.Println("portDiff", portDiffs)
	}

}
