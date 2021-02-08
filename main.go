package main

import (
	"os"

	"github.com/thomaspepio/rest-fizzbuzz/endpoint"
	"github.com/thomaspepio/rest-fizzbuzz/service"
)

const (
	// PortEnvVar : name of the environment variable to specify a port to listen to
	PortEnvVar = "PORT"

	// DefaultPort : default port the server will try to use
	DefaultPort = "8080"
)

func main() {
	port := os.Getenv(PortEnvVar)
	if port == "" {
		port = DefaultPort
	}

	urlCounter := map[service.FizzBuzzRequest]int{}
	router := endpoint.SetupRouter(port, urlCounter)
	router.Run(":" + port)
}
