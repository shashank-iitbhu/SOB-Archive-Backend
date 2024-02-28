package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/api", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "The API is running and healthy.",
		})
	})
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Println("Router was not able to run and map the requests")
		fmt.Println(err)
	}
}
