package rebalance

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

var API_URL = os.Getenv("SUPABASE_API_URL")
var API_KEY = os.Getenv("SUPABASE_API_KEY")

func getUsersByActiveStrategy(strategyId string) error {
	// gets all users with an active strategy of a specific strategyId from ActiveStrategies table
	// returns a list of userId, amount, and alpacaId
	fmt.Println(API_KEY, API_URL)
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initialize client", err)
		return err
	}
	data, count, err := client.From("ActiveStrategies").Select("*", "*", true).Eq("strategyId", strategyId).Execute()
	if err != nil {
		fmt.Println("cannot execute query", err)
		return err
	}
	fmt.Println(data, count)
	return nil
}
