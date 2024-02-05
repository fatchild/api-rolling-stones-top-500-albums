package logger

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLog(logToFile string) {
	if logToFile == "TRUE" {
		gin.DisableConsoleColor()
		dt := time.Now()
		f, _ := os.Create("access_" + dt.Format("2006_01_02") + ".log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}
