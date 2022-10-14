package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/initializers"
	"github.com/greazleay/vehicle-api/src/models"
)

func CreateVehicle(ctx *gin.Context) {

	var body struct {
		Model        string
		Category     string
		Year         int
		NumberOfSeat int
		Price        int
		Engine       *models.Engine
	}

	ctx.Bind(&body)

	maker := models.Vehicle{
		Model:        body.Model,
		Category:     body.Category,
		Year:         body.Year,
		NumberOfSeat: body.NumberOfSeat,
		Price:        body.Price,
		Engine:       *body.Engine,
	}

	result := initializers.DB.Create(&maker)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Maker Created",
		"maker":   maker,
	})
}
