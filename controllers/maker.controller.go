package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/initializers"
	"github.com/greazleay/vehicle-api/models"
)

func CreateMaker(ctx *gin.Context) {

	var body struct {
		Name string
	}

	ctx.Bind(&body)

	maker := models.Maker{Name: body.Name}

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

func GetAllMakers(ctx *gin.Context) {

	var makers []models.Maker
	initializers.DB.Find(&makers)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "All Makers",
		"data":    makers,
	})
}

func GetMakerByID(ctx *gin.Context) {

	id := ctx.Param("id")

	var maker models.Maker
	initializers.DB.First(&maker, "id = ?", id)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Maker",
		"data":    maker,
	})
}

func UpdateMaker(ctx *gin.Context) {

	var body struct {
		Name string
	}

	ctx.Bind(&body)

	id := ctx.Param("id")

	var maker models.Maker
	initializers.DB.First(&maker, "id = ?", id)

	initializers.DB.Model(&maker).Updates(models.Maker{Name: body.Name})

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Maker",
		"data":    maker,
	})

}

func DeleteMaker(ctx *gin.Context) {

	id := ctx.Param("id")

	initializers.DB.Delete(&models.Maker{}, "id = ?", id)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Maker Deleted",
	})
}
