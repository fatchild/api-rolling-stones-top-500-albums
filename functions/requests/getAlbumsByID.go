package requests

import (
	"database/sql"
	"net/http"

	"github.com/fatchild/api-rolling-stones-top-500-albums/database"
	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/utils"
	"github.com/gin-gonic/gin"
)

func GetAlbumByID(c *gin.Context) {
	var alb Album

	posString := utils.SanitisePositionParam(c.GetQuery("position"))
	yearString := utils.SanitiseYearParam(c.GetQuery("year"))

	queryString := "SELECT * FROM " + yearString + "_albums WHERE ID = ?"
	row := database.DB.QueryRow(queryString, posString)
	if err := row.Scan(&alb.Position, &alb.Artist, &alb.Album); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, alb)
	}
}
