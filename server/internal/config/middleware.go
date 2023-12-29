package config

import (
	"github.com/gin-gonic/gin"
	"sushi/internal/core/domain"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")

		if token == "" {
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
