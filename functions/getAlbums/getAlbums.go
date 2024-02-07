package getAlbums

import (
	"log"
	"net/http"

	"github.com/fatchild/api-rolling-stones-top-500-albums/database"
	"github.com/gin-gonic/gin"
)

type Album struct {
	Position int64
	Album    string
	Artist   string
}

func GetAlbumsJSON(c *gin.Context) {
	obj := database.ParseAlbumsJSON()
	c.IndentedJSON(http.StatusOK, obj)
}

func GetAlbums(c *gin.Context) {
	var albums []Album

	rows, err := database.DB.Query("SELECT * FROM 2024_albums")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.Position, &alb.Album, &alb.Artist); err != nil {
			log.Fatal(err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, albums)
}
