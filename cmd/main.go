package main

import (
	"billi-md/internal/handler"
	"billi-md/internal/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init logger
	logger.InitLogger()
	defer logger.CloseLogger()

	logger.Logger.Info("Starting server...")

	router := gin.Default()

	// set up a new Handler
	handler.NewHandler(&handler.Config{R: router})

	// Run the Gin server on port 8080
	if err := router.Run(":8080"); err != nil {
		logger.Logger.Fatalw("Failed to start server", "error", err)
	}

}
