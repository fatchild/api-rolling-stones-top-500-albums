package logger

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func GinLog(logToFile bool) {
	if logToFile {
		gin.DisableConsoleColor()
		dt := time.Now()
		f, _ := os.Create("var/log/access_" + dt.Format("2006_01_02") + ".log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}
