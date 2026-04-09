package main

import (
	"log"

	"local.exchange-demo/exchange-core-go/bootstrap"
)

func main() {
	if err := bootstrap.RunReplayTool(); err != nil {
		log.Fatal(err)
	}
}
