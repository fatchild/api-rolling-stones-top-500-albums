package main

import (
	"github.com/fatchild/api-rolling-stones-top-500-albums/logger"
	"github.com/fatchild/api-rolling-stones-top-500-albums/router"
)

const (
	version           = "0.0.1-alpha"
	routingServiceURL = "localhost:"
	port              = "8000"
	logToFile         = true
	releaseMode       = false
)

func main() {
	logger.GinLog(logToFile)
	router.Router(version, routingServiceURL, port, releaseMode)
}
