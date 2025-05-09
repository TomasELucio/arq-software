package main

import (
	"API-VINILOS/config"
	"API-VINILOS/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate()

	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.PostAlbums)
	router.GET("/albums/:id", handlers.GetAlbumsById)
	router.PUT("/albums/:id", handlers.PutAlbumById)
	router.DELETE("/albums/:id", handlers.DeleteById)
	router.Run()
}
