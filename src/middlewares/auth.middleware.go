package middlewares

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/exceptions"
	"github.com/greazleay/vehicle-api/src/services/auth"
)

func Auth() gin.HandlerFunc {

	return func(context *gin.Context) {

		bearerToken := context.GetHeader("Authorization")
		if bearerToken == "" {
			exceptions.HandleBadRequestException(context, errors.New("bearer token is required"))
			return
		}

		accessToken := strings.Split(bearerToken, "Bearer ")[1]
		err := auth.ValidateToken(accessToken)
		if err != nil {

			exceptions.HandleUnauthorizedException(context, "Unauthorized")
			return
		}
		context.Next()
	}
}
