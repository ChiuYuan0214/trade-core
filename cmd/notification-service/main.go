package main

import (
	"log"

	"exchange-demo/internal/bootstrap"
)

func main() {
	if err := bootstrap.RunProcess("notification-service"); err != nil {
		log.Fatal(err)
	}
}
