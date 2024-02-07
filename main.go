package main

import (
	"embed"
	"os"

	"github.com/fatchild/api-rolling-stones-top-500-albums/database"
	"github.com/fatchild/api-rolling-stones-top-500-albums/environment"
	"github.com/fatchild/api-rolling-stones-top-500-albums/logger"
	"github.com/fatchild/api-rolling-stones-top-500-albums/router"
)

const (
	Version = "0.0.1-alpha"
)

var resources embed.FS

func main() {
	environment.Load("./")
	database.Connect()
	logger.GinLog(os.Getenv("LOG_TO_FILE"))
	router.Router(Version, os.Getenv("URL"), os.Getenv("PORT"), os.Getenv("RELEASE_MODE"))
}
