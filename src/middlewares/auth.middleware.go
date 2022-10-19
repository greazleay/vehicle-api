package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/services/auth"
)

func Auth() gin.HandlerFunc {

	return func(context *gin.Context) {

		bearerToken := context.GetHeader("Authorization")
		if bearerToken == "" {

			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"statusText": "failed",
				"statusCode": 400,
				"errorType":  "BadRequestException",
				"error":      "Access Token is Required",
			})
			return
		}

		accessToken := strings.Split(bearerToken, "Bearer ")[1]
		err := auth.ValidateToken(accessToken)
		if err != nil {

			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"statusText": "failed",
				"statusCode": 401,
				"errorType":  "UnauthorizedException",
				"error":      "Unauthorized",
			})
			return
		}
		context.Next()
	}
}
