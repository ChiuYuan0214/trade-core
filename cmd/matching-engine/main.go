package main

import (
	"log"

	"exchange-demo/internal/bootstrap"
)

func main() {
	if err := bootstrap.RunProcess("matching-engine"); err != nil {
		log.Fatal(err)
	}
}
