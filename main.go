package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	rebalance "github.com/willgolden5/blackbox-rebalance-service/rebalance"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	router := gin.Default()
	apiKey := os.Getenv("API_KEY")

	router.GET("/rebalance", func(ctx *gin.Context) {
		requestKey := ctx.GetHeader("X-API-KEY")
		if requestKey != apiKey {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			return
		}

		activeStrategyId := ctx.Query("activeStrategyId")

		go rebalance.RebalanceUserPortfolio(activeStrategyId)
		fmt.Println("rebalance initiated")
		ctx.JSON(200, gin.H{
			"data": "rebalance initiated",
		})
	})

	router.Run(":8080")
}
