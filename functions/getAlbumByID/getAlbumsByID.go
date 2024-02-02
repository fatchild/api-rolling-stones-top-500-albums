package getAlbumByID

import (
	"net/http"
	"strconv"

	"github.com/fatchild/api-rolling-stones-top-500-albums/database"
	"github.com/gin-gonic/gin"
)

func GetAlbumByID(c *gin.Context) {
	obj := database.ParseAlbumsJSON()

	idString := c.Param("position")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	for _, a := range obj {
		if a.Position == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
