package main

import (
	"github.com/fatchild/api-rolling-stones-top-500-albums/database"
	"github.com/fatchild/api-rolling-stones-top-500-albums/environment"
	"github.com/fatchild/api-rolling-stones-top-500-albums/logger"
	"github.com/fatchild/api-rolling-stones-top-500-albums/router"
)

func main() {
	environment.Load("./")
	database.Connect()
	logger.GinLog()
	router.Router()
}
