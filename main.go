package main

import (
	"example/Auction-Service/database"

	"github.com/gin-gonic/gin"
)

var serverURI = "localhost:8081"

func main() {

	router := gin.Default()

	router.GET("/auctions/:id/winning-bid", database.GetWinningBid)

	router.Run(serverURI)
}
