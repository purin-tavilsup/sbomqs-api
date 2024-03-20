package main

import (
	"github.com/gin-gonic/gin"
	"github.com/purin-tavilsup/sbomqs-api/internal/handler"
	"log"
)

func main() {
	log.Println("Starting server listening on port 5050...")

	router := gin.Default()

	handler.AddHandler(&handler.Config{
		Route: router,
	})

	err := router.Run(":5050")

	if err != nil {
		log.Println("An error occurred while running the server.")
		return
	}

	log.Println("Shutting down server...")
}
