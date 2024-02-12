package router

import (
	"log"
	"os"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/pages"
	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/requests"
	"github.com/gin-gonic/gin"
)

func Router() {
	if os.Getenv("RELEASE_MODE") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("favicon.png", "./resources/favicon.png")
	router.GET("/", pages.Index)
	router.GET("/getAlbumList", requests.GetAlbums)
	router.GET("/getAlbumAtPosition", requests.GetAlbumByID)
	err := router.Run(os.Getenv("URL") + ":" + os.Getenv("PORT"))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
	log.Println("listening on", os.Getenv("PORT"))
}
