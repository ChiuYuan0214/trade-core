package main

import (
	"log"

	"exchange-demo/internal/bootstrap"
)

func main() {
	if err := bootstrap.RunProcess("rest-gateway"); err != nil {
		log.Fatal(err)
	}
}
