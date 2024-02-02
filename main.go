package main

import (
	"github.com/fatchild/api-rolling-stones-top-500-albums/logger"
	"github.com/fatchild/api-rolling-stones-top-500-albums/router"
)

const (
	version			  = "0.0.1-alpha"
	routingServiceURL = "http://localhost"
	port              = "8080"
	logToFile	      = true
)

func main() {
	logger.GinLog(logToFile)
	router.Router(version, routingServiceURL, port)
}
