package pages

import (
	"net/http"
	"os"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/consts"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	url := os.Getenv("URL") + ":" + os.Getenv("PORT")
	if os.Getenv("PUBLIC_URL") != "" {
		url = os.Getenv("PUBLIC_URL")
	}

	c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
		"version": consts.VERSION,
		"url":     url,
	})
}
