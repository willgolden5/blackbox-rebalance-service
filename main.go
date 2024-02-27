package main

import (
	"github.com/gin-gonic/gin"
	rebalance "github.com/willgolden5/blackbox-rebalance-service/rebalance"
)

func main() {
	router := gin.Default()

	router.GET("/rebalance", func(c *gin.Context) {
		rebalance.RebalanceUserPortfolios("test")
		c.JSON(200, gin.H{
			"message": "hit",
		})
	})

	router.Run(":8080")
}
