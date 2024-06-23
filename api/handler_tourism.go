package api

import (
	"e/database"
	"e/model"
	"fmt"
	gin "github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TourismHandler struct {
	conn *pgx.Conn
}

func NewTourismHandler(conn *pgx.Conn) *TourismHandler {
	return &TourismHandler{conn: conn}
}

func (h TourismHandler) Handler() *gin.Engine {
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

// @Summary Create Tourism.
// @Description Notes:<br>Leave the "relations" sub-entity empty.
// @Tags Tourism
// @Produce json
// @Param request body model.Tourism true "Tourism"
// @Success 200 {object} api.APIError
// @Failure 400 {object} api.APIError
// @Failure 500 {object} api.APIError
// @Router /createTourism [post]
func (h TourismHandler) Post(ctx *gin.Context) {
	var payload model.Tourism
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	payload.Id = uuid.New()
	tourismData, err := database.PostTourism(h.conn, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		return
	}

	ctx.JSON(http.StatusOK, tourismData)
}

// @Summary Read Tourism.
// @Tags Tourism
// @Produce json
// @Param id path string true "Tourism ID"
// @Success 200 {object} model.Tourism
// @Failure 404 {object} api.APIError
// @Failure 500 {object} api.APIError
// @Router /api/tags/readTourism/{id} [get]
func (h TourismHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	// Check if the data with ID exists in the database
	tourismData, err := database.GetTourism(h.conn, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	ctx.JSON(http.StatusOK, tourismData)
}

// @Summary Update Tourism.
// @Tags Tourism
// @Produce json
// @Param id path string false "Toursim"
// @Param request body model.Tourism true "Tourism"
// @Success 204
// @Failure 400 {object} api.APIError
// @Failure 404
// @Failure 500 {object} api.APIError
// @Router /updateTourism/{id} [put]
func (h TourismHandler) Put(ctx *gin.Context) {
	id := ctx.Param("id")

	// Check if the data with the specified id exists
	var payload model.Tourism
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if id != payload.Id.String() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Different Id"})
		return
	}
	err := database.PutTourism(h.conn, id, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data in the database"})
		return
	}

	// Return the updated data to the client
	ctx.JSON(http.StatusOK, payload)
}

// @Summary Delete Tourism.
// @Tags Tourism
// @Produce json
// @Param id path string false "Toursim"
// @Success 204
// @Failure 404
// @Failure 500 {object} api.APIError
// @Router /deleteTourism/{id} [delete]
func (h TourismHandler) Del(ctx *gin.Context) {
	fmt.Println("verrosa")
	id := ctx.Param("id")

	// Check id the data with ID there is database
	if err := database.DelTourism(h.conn, id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}
	ctx.Status(http.StatusOK)
}
