package getAlbumByID

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/fatchild/api-rolling-stones-top-500-albums/database"
	"github.com/gin-gonic/gin"
)

type Album struct {
	Position int64
	Album    string
	Artist   string
}

func GetAlbumByIDJSON(c *gin.Context) {
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

func GetAlbumByID(c *gin.Context) {
	var alb Album

	idString := c.Param("position")

	row := database.DB.QueryRow("SELECT * FROM 2024_albums WHERE ID = ?", idString)
	if err := row.Scan(&alb.Position, &alb.Album, &alb.Artist); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, alb)
	}
}
