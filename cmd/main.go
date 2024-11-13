package main

import (
	"billi-md/internal/handler"
	"billi-md/internal/logger"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Init logger
	logger.InitLogger()
	defer logger.CloseLogger()

	connStr := "postgres://knassar1:dbTest123@localhost:5432/billi-db?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Logger.Fatalw("Failed to connect to DB", "error", err)
	}
	defer db.Close()

	logger.Logger.Info("Starting server...")
	router := gin.Default()

	// set up a new Handler
	handler.NewHandler(&handler.Config{R: router})

	// Run the Gin server on port 8080
	if err := router.Run(":8080"); err != nil {
		logger.Logger.Fatalw("Failed to start server", "error", err)
	}

}
