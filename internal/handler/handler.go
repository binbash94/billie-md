package handler

import (
	"net/http"

	"billi-md/internal/services"

	"github.com/gin-gonic/gin"
)

// TODO: add the DB dependency
// Config which holds any dependies that the Handler services requires
type Config struct {
	R *gin.Engine
}

// Used to store required services for handler to function.
// These will later be linked to routers when http requests are called.
// TODO: add the DB dependency, this db dependency will be set from the config level
type Handler struct {
	service *services.DataService
}

// This function initializes a new Handler with the provided configuration
// It accepts a Config parameter which provides the router and any other services that Handler may need
func NewHandler(config *Config) {
	// create instance of handler which will later have injected services
	// TODO: inject the db dependency in the handler to pass onto the service layer
	handler := &Handler{}

	config.R.GET("/api/", handler.generateBillingCode)
}

// Placeholder implementing the Handler type to get the require billing code.
// TODO: call the service layer method to get the appropriate billing code from DB
func (handler *Handler) generateBillingCode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "its GenerateBillingCode",
	})
}
