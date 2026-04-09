package main

import (
	"log"

	"local.exchange-demo/exchange-core-go/bootstrap"
)

func main() {
	if err := bootstrap.RunWSGateway(); err != nil {
		log.Fatal(err)
	}
}
