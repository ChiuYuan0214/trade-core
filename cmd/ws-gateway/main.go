package main

import (
	"log"

	"exchange-demo/internal/bootstrap"
)

func main() {
	if err := bootstrap.RunProcess("ws-gateway"); err != nil {
		log.Fatal(err)
	}
}
