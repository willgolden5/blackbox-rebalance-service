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

	router.GET("/rebalance", func(c *gin.Context) {
		rebalance.RebalanceUserPortfolios("congress_buys")
		c.JSON(200, gin.H{
			"message": "hit",
		})
	})

	router.Run(":8080")
}
