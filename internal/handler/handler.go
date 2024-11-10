package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: add the DB dependency
// Config which holds any dependies that the Handler services requires
type Config struct {
	R *gin.Engine
}

// Used to store required services for handler to function.
// These will later be linked to routers when http requests are called.
type Handler struct{}

// This function initializes a new Handler with the provided configuration
// It accepts a Config parameter which provides the router and any other services that Handler may need
func NewHandler(config *Config) {
	// create instance of handler which will later have injected services
	handler := &Handler{}

	config.R.GET("/api/", handler.GenerateBillingCode)
}

// Placeholder implementing the Handler type to get the require billing code.
func (handler *Handler) GenerateBillingCode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "its GenerateBillingCode",
	})
}
