package rebalance

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

func getUsersWithActiveStrategy(strategyId string) ([]StrategyUser, error) {
	// gets all users with an active strategy of a specific strategyId from ActiveStrategies table
	// returns a list of userId, amount, and alpacaId
	var API_URL = os.Getenv("SUPABASE_API_URL")
	var API_KEY = os.Getenv("SUPABASE_API_KEY")
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initialize client", err)
		return nil, err
	}

	data, count, err := client.From("ActiveStrategies").Select("*", "exact", false).Eq("strategyId", strategyId).Execute()
	if err != nil {
		fmt.Println("cannot execute query", err)
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	var result []StrategyUser
	jsonErr := json.Unmarshal(data, &result)
	if jsonErr != nil {
		fmt.Println("error:", jsonErr)
	}
	return result, nil
}
