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

	yearString := "2023"
	yearQuery, yearGiven := c.GetQuery("year")
	if yearGiven && (yearQuery == "2023" || yearQuery == "2020" || yearQuery == "2012" || yearQuery == "2003") {
		yearString = yearQuery
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "year not valid. Valid years are 2003, 2012, 2020 & 2023."})
		return
	}

	queryString := "SELECT * FROM " + yearString + "_albums"
	rows, err := database.DB.Query(queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.Position, &alb.Artist, &alb.Album); err != nil {
			log.Fatal(err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, albums)
}
