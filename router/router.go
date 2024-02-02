package router

import (
	"net/http"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/getAlbumByID"
	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/getAlbums"
	"github.com/gin-gonic/gin"
)

func Router(version string, routingServiceURL string, port string) {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"version": version,
		})
	})
	router.GET("/getAlbumList", getAlbums.GetAlbums)
	router.GET("/getAlbumAtPosition/:position", getAlbumByID.GetAlbumByID)

	router.Run("localhost:8080")

	err := router.Run(routingServiceURL + port)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}
