package main

import (
	"log"

	"exchange-demo/internal/bootstrap"
)

func main() {
	if err := bootstrap.RunProcess("replay-tool"); err != nil {
		log.Fatal(err)
	}
}
