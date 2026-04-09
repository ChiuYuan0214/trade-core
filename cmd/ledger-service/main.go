package main

import (
	"log"

	"exchange-demo/internal/bootstrap"
)

func main() {
	if err := bootstrap.RunProcess("ledger-service"); err != nil {
		log.Fatal(err)
	}
}
