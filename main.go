package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type returnData struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

var possibleReturns = []returnData{
	{ID: "0", Message: "This is a test"},
	{ID: "1", Message: "This is another test"},
	{ID: "2", Message: "Yet another test la"},
}

func getResult(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, possibleReturns)
}

func main() {
	router := gin.Default()
	router.GET("/results", getResult)

	router.Run("0.0.0.0:16489")
}
