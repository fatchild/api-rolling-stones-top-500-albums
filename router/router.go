package router

import (
	"log"
	"net/http"
	"os"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/getAlbumByID"
	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/getAlbums"
	"github.com/gin-gonic/gin"
)

func Router(version string, routingServiceURL string, port string, releaseMode string) {
	if releaseMode == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	URL_PORT := routingServiceURL + ":" + port

	url := "http://" + URL_PORT
	if os.Getenv("PUBLIC_URL") != "" {
		url = os.Getenv("PUBLIC_URL")
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
			"version": version,
			"url":     url,
		})
	})
	//router.GET("/getAlbumList", getAlbums.GetAlbumsJSON)
	router.GET("/getAlbumList", getAlbums.GetAlbums)
	//router.GET("/getAlbumAtPosition/:position", getAlbumByID.GetAlbumByIDJSON)
	router.GET("/getAlbumAtPosition", getAlbumByID.GetAlbumByID)

	err := router.Run(URL_PORT)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
	log.Println("listening on", port)
}
