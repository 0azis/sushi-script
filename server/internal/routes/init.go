package routes

import (
	"github.com/gin-gonic/gin"
	"sushi/internal/store"
)

func InitRoutes(router *gin.RouterGroup, db store.Store, m gin.HandlerFunc) {
	userRoutes(router, db, m)
}
