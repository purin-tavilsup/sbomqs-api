package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/_health", getHealth)

	err := server.Run(":5050")

	if err != nil {
		return
	}
}

func getHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Healthy"})
}
