package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h TourismHandler) Router() *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")

	// Add a handler for the create endpoint
	baseRouter.POST("/createTourism", h.Post)
	// Add a handler for the read endpoint
	baseRouter.GET("/readTourism/:id", h.Get)
	// Add a handler for the update endpoint
	baseRouter.PUT("/updateTourism/:id", h.Put)
	// Add a handler for the delete endpoint
	baseRouter.DELETE("/deleteTourism/:id", h.Del)

	return router
}
