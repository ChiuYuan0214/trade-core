package main

import (
	"log"

	"exchange-demo/internal/bootstrap"
)

func main() {
	if err := bootstrap.RunProcess("market-data-service"); err != nil {
		log.Fatal(err)
	}
}
