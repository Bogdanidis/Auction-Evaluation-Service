package main

import (
	"bytes"
	"encoding/json"
	"example/Auction-Service/database"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetWinningBid(t *testing.T) {
	router := SetUpRouter()
	auctionId := "62c3b3350ff66c5091bf94bc"
	router.GET("/auctions/:id/winning-bid", database.GetWinningBid)
	req, _ := http.NewRequest("GET", "/auctions/"+auctionId+"/winning-bid", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var bid database.Bid
	json.Unmarshal(w.Body.Bytes(), &bid)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("GET", "/auctions/1/winning-bid", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
