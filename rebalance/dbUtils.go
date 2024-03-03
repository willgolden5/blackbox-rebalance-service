package rebalance

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

func getActiveStrategyListing(activeStrategyId string) (*ActiveStrategy, error) {
	var API_URL = os.Getenv("SUPABASE_API_URL")
	var API_KEY = os.Getenv("SUPABASE_API_KEY")
	// gets all users with an active strategy of a specific strategyId from ActiveStrategies table
	// returns a list of userId, amount, and alpacaId
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initialize client", err)
		return nil, err
	}

	data, count, err := client.From("ActiveStrategies").Select("*", "exact", false).Eq("id", activeStrategyId).Execute()
	if err != nil {
		fmt.Println("cannot execute query", err)
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	var results []ActiveStrategy
	jsonErr := json.Unmarshal(data, &results)
	if jsonErr != nil {
		fmt.Println("error:", jsonErr)
	}
	if len(results) == 0 {
		return nil, nil
	}

	return &results[0], nil
}

func getStrategyHoldings(table string) ([]Position, error) {
	// get the holdings of the strategy specified from the strategy table
	// returns a list of equities and their respective weights
	var API_URL = os.Getenv("SUPABASE_API_URL")
	var API_KEY = os.Getenv("SUPABASE_API_KEY")
	// gets all users with an active strategy of a specific strategyId from ActiveStrategies table
	// returns a list of userId, amount, and alpacaId
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initialize client", err)
		return nil, err
	}
	data, count, err := client.From("CongressBuys").Select("ticker,navPercentage", "exact", false).Execute()
	if err != nil {
		fmt.Println("cannot execute query", err)
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	var results []Position
	jsonErr := json.Unmarshal(data, &results)
	if jsonErr != nil {
		fmt.Println("error:", jsonErr)
	}
	if len(results) == 0 {
		return nil, nil
	}

	return results, nil
}
