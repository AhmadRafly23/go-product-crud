package middleware

import (
	"net/http"
	"strings"

	"github.com/AhmadRafly23/go-product-crud/helper"
	"github.com/gin-gonic/gin"
)

func BearerAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")
		// {Authorization: Bearer jwt_token}
		// get the encoded string
		splitToken := strings.Split(headerAuth, " ")
		if len(splitToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization header",
			})
			return
		}

		// check basic
		if splitToken[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization method",
			})
			return
		}
		// validate jwt
		valid, email, _ := helper.ValidateUserJWT(splitToken[1])
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "malformed token",
			})
			return
		}

		ctx.Set("user", email)

		ctx.Next()
	}
}
