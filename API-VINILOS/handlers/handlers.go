package handlers

import (
	"API-VINILOS/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "The Number of the Beast", Artist: "Iron Maiden", Year: 1982, Price: 25.19},
	{ID: "2", Title: "Youthanasia", Artist: "Medadeth", Year: 1994, Price: 13.65},
	{ID: "3", Title: "Master of Puppets", Artist: "Metallica", Year: 1986, Price: 20.97},
}

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)

}

func PostAlbums(ctx *gin.Context) {
	var newAlbum models.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumsById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"messge": "album no encontrado"})
}

func PutAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyAlbum models.Album
	if err := ctx.BindJSON(&modifyAlbum); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "datos incorrectos"})
		return
	}
	for i, a := range albums {
		if a.ID == id {
			albums[i] = modifyAlbum
			albums[i].ID = id
			ctx.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
}

func DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "album eliminado"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})

}
