package config

import (
	"github.com/gin-gonic/gin"
	"sushi/internal/routes"
	"sushi/internal/store"
)

func InitHttpServer() {

	app := gin.Default()

	db := store.InitDB()
	db.Connect()
	defer db.Close()

	api := app.Group("/api")

	routes.InitRoutes(api, *db, JWTMiddleware())

	app.Run()
}
