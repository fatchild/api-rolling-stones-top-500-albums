package logger

import (
	"io"
	"os"
	"time"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/utils"
	"github.com/gin-gonic/gin"
)

func GinLog() {
	if os.Getenv("LOG_TO_FILE") == "TRUE" {
		gin.DisableConsoleColor()
		dt := time.Now()
		utils.CreateDirectory(os.Getenv("LOG_FILE_PATH"))
		f, _ := os.Create(os.Getenv("LOG_FILE_PATH") + "access_" + dt.Format("2006_01_02") + ".log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}
