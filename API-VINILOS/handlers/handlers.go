package handlers

import (
	"API-VINILOS/config"
	"API-VINILOS/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(ctx *gin.Context) {
	var albums []models.Album
	config.DB.Find(&albums)
	ctx.IndentedJSON(http.StatusOK, albums)

}

func PostAlbums(ctx *gin.Context) {
	var newAlbum models.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	config.DB.Create(&newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumsById(ctx *gin.Context) {
	id := ctx.Param("id")
	var album models.Album

	result := config.DB.First(&album, "id = ?", id)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"messge": "album no encontrado"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, album)
}

func PutAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyAlbum models.Album
	var album models.Album

	result := config.DB.First(&album, "id = ?", id)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"messge": "album no encontrado"})
		return
	}

	if err := ctx.BindJSON(&modifyAlbum); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "datos incorrectos"})
		return
	}

	album.Title = modifyAlbum.Title
	album.Title = modifyAlbum.Title
	album.Title = modifyAlbum.Title
	album.Price = modifyAlbum.Price
	album.Year = modifyAlbum.Year

	config.DB.Save(&album)
	ctx.IndentedJSON(http.StatusOK, album)
}

func DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	result := config.DB.Delete(&models.Album{}, "id = ?", id)

	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "album eliminado"})

}
