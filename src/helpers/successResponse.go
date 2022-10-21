package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleOkResponse(context *gin.Context, message string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    message,
		"data":       data,
	})
}

func HandleCreatedResponse(context *gin.Context, message string, data interface{}) {
	context.JSON(http.StatusCreated, gin.H{
		"statusText": "success",
		"statusCode": 201,
		"message":    message,
		"data":       data,
	})
}
