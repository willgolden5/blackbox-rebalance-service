package rebalance

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/shopspring/decimal"
)

func calculateNAV(positions []alpaca.Position, buyingPower decimal.Decimal) []Position {
	var navs []Position
	// Convert decimal.Decimal to float64
	for _, position := range positions {
		nav := Position{
			Ticker: position.Symbol,
			NAV:    position.MarketValue.Div(buyingPower).Round(2),
		}
		navs = append(navs, nav)
	}
	return navs
}

func convertToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func compareHoldings(portfolioData []Position, quiverData []Position) []PortDiff {
	portfolioMap := make(map[string]decimal.Decimal)
	for _, item := range portfolioData {
		portfolioMap[item.Ticker] = item.NAV
	}

	var differences []PortDiff
	for _, quiverItem := range quiverData {
		portfolioPercentOfNav, exists := portfolioMap[quiverItem.Ticker]
		if !exists {
			differences = append(differences, PortDiff{
				Ticker:       quiverItem.Ticker,
				PercentOfNav: quiverItem.NAV,
				Difference:   "Not in portfolio",
			})
		} else if !portfolioPercentOfNav.Equal(quiverItem.NAV) {
			navDifference := quiverItem.NAV.Sub(portfolioPercentOfNav)
			differences = append(differences, PortDiff{
				Ticker:       quiverItem.Ticker,
				PercentOfNav: quiverItem.NAV,
				Difference:   navDifference,
			})
		}
	}

	sort.Slice(differences, func(i, j int) bool {
		diffI, okI := differences[i].Difference.(decimal.Decimal)
		diffJ, okJ := differences[j].Difference.(decimal.Decimal)
		if okI && okJ {
			return diffI.LessThan(diffJ)
		} else if okI {
			return true
		} else if okJ {
			return false
		} else {
			return false
		}
	})

	return differences
}

func awaitMarketOpen(client *alpaca.Client) {
	for {
		clock, err := client.GetClock()
		if err != nil {
			fmt.Println("error getting clock", err)
			time.Sleep(time.Minute)
			continue
		}

		if clock.IsOpen {
			break
		} else {
			openTime := clock.NextOpen
			currTime := clock.Timestamp
			timeToOpen := openTime.Sub(currTime)
			fmt.Printf("%s until market open.\n", timeToOpen.String())
			time.Sleep(time.Minute)
		}
	}
}
