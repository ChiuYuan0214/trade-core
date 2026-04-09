package main

import (
	"log"

	"local.exchange-demo/exchange-core-go/bootstrap"
)

func main() {
	if err := bootstrap.RunLedgerService(); err != nil {
		log.Fatal(err)
	}
}
