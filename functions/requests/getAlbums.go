package requests

import (
	"log"
	"net/http"

	"github.com/fatchild/api-rolling-stones-top-500-albums/database"
	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/utils"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []Album

	yearString := utils.SanitiseYear(c.GetQuery("year"))

	queryString := "SELECT * FROM " + yearString + "_albums"
	rows, err := database.DB.Query(queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.Position, &alb.Artist, &alb.Album); err != nil {
			log.Fatal(err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Database error."})
	} else {
		c.IndentedJSON(http.StatusOK, albums)
	}
}
