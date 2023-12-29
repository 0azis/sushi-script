package routes

import (
	"github.com/gin-gonic/gin"
	"sushi/internal/adapters/http"
	"sushi/internal/core/services"
	"sushi/internal/store"
)

func userRoutes(router *gin.RouterGroup, db store.Store, m gin.HandlerFunc) {
	user := router.Group("/user")

	service := services.NewUserService(&db)
	adapters := http.NewUserAdapters(&service)

	user.POST("/login", adapters.Login)
	user.PATCH("", m, adapters.UpdateProfile)
	user.GET("", m, adapters.GetMyProfile)
}
