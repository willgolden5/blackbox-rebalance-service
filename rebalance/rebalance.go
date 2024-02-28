package rebalance

import (
	"fmt"
)

// call this in a goRoutine
func RebalanceUserPortfolios(strategyId string) {
	// FIRST:
	// get users from ActiveStrategies table where active strategyId = strategyId
	data, err := getUsersWithActiveStrategy(strategyId)
	if err != nil {
		fmt.Println("error getting users by active strategy", err)
		return
	}
	if data == nil {
		fmt.Println("no users with active strategy", strategyId)
		return
	}
	fmt.Println(data)
	// for each user:
	// get their alpaca account =>
	// get the active positions from their alpaca account =>
	// create object with alpacaId, accountId, and positions object with nav calculated for each position
	// relative to their tradeable balance in their alpaca account

	// THEN:
	// get the holdings of the strategy specified from the strategy table =>
	// compare the strategy holdings to the users current portfolio holdings and output a portDiff struct =>
	// if (portDiff is not empty) => sell all equities that need to be sold THEN
	// buy all equities that need to be bought =>
	// update the user's portfolio in the ActiveStrategies table with rebalance date
	fmt.Println(strategyId)
}
