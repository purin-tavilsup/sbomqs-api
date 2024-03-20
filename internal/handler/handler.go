package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/purin-tavilsup/sbomqs-api/internal/models"
	"net/http"
	"os"
)

type Handler struct{}

type Config struct {
	Route *gin.Engine
}

// AddHandler initializes the handler with required injected services along with http routes
func AddHandler(config *Config) {
	handler := &Handler{}
	group := config.Route.Group("/api/v1/sbomqs")

	group.GET("/_health", handler.getHealth)
	group.POST("/score", handler.evaluateSbom)
}

// getHealth returns health status of the service
func (handler *Handler) getHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Healthy",
	})
}

// evaluateSbom evaluates quality of sbom and returns a score
func (handler *Handler) evaluateSbom(context *gin.Context) {
	var request models.EvaluationSbomRequest

	if err := context.BindJSON(&request); err != nil {
		return
	}

	// Write SBOM to a file
	fileName := uuid.NewString()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	numberOfByte, err := file.WriteString(request.Sbom)
	if err != nil {
		fmt.Println(err)
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	fmt.Println(numberOfByte, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Evaluate SBOM file

	// Delete the file

	context.IndentedJSON(http.StatusOK, request.Sbom)
}
