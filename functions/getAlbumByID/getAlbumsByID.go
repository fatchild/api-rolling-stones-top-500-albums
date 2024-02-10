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

	posString, posGiven := c.GetQuery("position")
	pos, err := strconv.Atoi(posString)
	if err != nil || !posGiven {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "position not valid. Number between 1 and 500 inclusive."})
		return
	}

	for _, a := range obj {
		if a.Position == pos {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func GetAlbumByID(c *gin.Context) {
	var alb Album

	posString, posGiven := c.GetQuery("position")
	pos, err := strconv.Atoi(posString)
	if err != nil || !posGiven || pos < 1 || pos > 500 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "position not valid. Number between 1 and 500 inclusive."})
		return
	}

	yearString := "2023"
	yearQuery, yearGiven := c.GetQuery("year")
	if yearGiven && (yearQuery == "2023" || yearQuery == "2020" || yearQuery == "2012" || yearQuery == "2003") {
		yearString = yearQuery
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "year not valid. Valid years are 2003, 2012, 2020 & 2023."})
		return
	}

	queryString := "SELECT * FROM " + yearString + "_albums WHERE ID = ?"
	row := database.DB.QueryRow(queryString, pos)
	if err := row.Scan(&alb.Position, &alb.Artist, &alb.Album); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, alb)
	}
}
