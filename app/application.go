package app

import (
	"log"

	"github.com/edilbertloquine/go-microservices/users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
