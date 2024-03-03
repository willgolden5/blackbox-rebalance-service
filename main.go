package main

import (
	"fmt"

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

	router.GET("/rebalance", func(ctx *gin.Context) {
		activeStrategyId := ctx.Query("activeStrategyId")

		go rebalance.RebalanceUserPortfolio(activeStrategyId)
		ctx.JSON(200, gin.H{
			"message": "hit",
		})
	})

	router.Run(":8080")
}
