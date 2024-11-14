package handler

import (
	"database/sql"
	"net/http"

	"billi-md/internal/models"
	"billi-md/internal/services"

	"github.com/gin-gonic/gin"
)

// Config which holds any dependies that the Handler services requires
type Config struct {
	R  *gin.Engine
	db *sql.DB
}

// Used to store required services for handler to function
type Handler struct {
	service *services.DataService
}

// This function initializes a new Handler with the provided configuration
func NewHandler(config *Config) {
	// configure the handler with the dependencies required to call the service layer
	handler := &Handler{service: services.NewDataService(config.db)}

	config.R.GET("/api/", handler.generateBillingCode)
}

// Placeholder implementing the Handler type to get the require billing code.
func (handler *Handler) generateBillingCode(ctx *gin.Context) {

	// define the qualifier model that we will bind our user POST data to
	var Qualifiers models.Qualifiers

	// call the service layer method
	responseCode := handler.service.GetBillingCode(&Qualifiers)

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode": responseCode,
	})
}
