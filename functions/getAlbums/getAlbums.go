package getAlbums

import (
	"net/http"

	"github.com/fatchild/api-rolling-stones-top-500-albums/database"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	obj := database.ParseAlbumsJSON()
	c.IndentedJSON(http.StatusOK, obj)

	database.Connect()
}
