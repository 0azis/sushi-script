package config

import (
	"github.com/gin-gonic/gin"
	"strings"
	"sushi/helpers"
	"sushi/internal/core/domain"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]

		if len(header) == 0 {
			c.AbortWithStatusJSON(401, domain.HttpResponse{
				Error: domain.Error{
					Status:  401,
					Message: "Unauthorized",
				},
			})
			return
		}

		token := strings.Split(header[0], " ")[1]

		if _, err := helpers.GetIdentity(token); err != nil {
			c.AbortWithStatusJSON(401, domain.HttpResponse{
				Error: domain.Error{
					Status:  401,
					Message: "Unauthorized",
				},
			})
			return
		}

	}

}
