package main

import (
	"github.com/thomaspepio/rest-fizzbuzz/endpoint"
	"github.com/thomaspepio/rest-fizzbuzz/service"
)

func main() {
	urlCounter := map[service.FizzBuzzRequest]int{}
	router := endpoint.SetupRouter(urlCounter)
	router.Run(":8080")
}
