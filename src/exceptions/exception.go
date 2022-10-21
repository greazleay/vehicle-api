package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBadRequestException(context *gin.Context, err error) {
	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"statusText": "failure",
		"statusCode": 400,
		"errorType":  "BadRequestException",
		"error":      err.Error(),
	})
}

func HandleConflictException(context *gin.Context, errText string) {
	context.AbortWithStatusJSON(http.StatusConflict, gin.H{
		"statusText": "failed",
		"statusCode": 409,
		"errorType":  "ConflictException",
		"error":      errText,
	})
}

func HandleInternalServerException(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"statusText": "failed",
		"statusCode": 500,
		"errorType":  "InternalServerErrorException",
		"error":      "Something Went Wrong",
	})
}

func HandleNotFoundException(context *gin.Context, err error) {
	context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"statusText": "failure",
		"statusCode": 404,
		"errorType":  "NotFoundException",
		"error":      err,
	})
}

func HandleUnauthorizedException(context *gin.Context, errText string) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"statusText": "failed",
		"statusCode": 401,
		"errorType":  "UnauthorizedException",
		"error":      errText,
	})
}

func HandleValidationException(context *gin.Context, err error) {
	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"statusText": "failure",
		"statusCode": 400,
		"errorType":  "ValidationException",
		"error":      err.Error(),
	})
}
