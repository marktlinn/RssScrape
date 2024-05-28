package main

import (
	"log"
	"os"

	"github.com/marktlinn/RssScrape/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("failed to find port: PORT env variable must be set")
	}

	server.Server(port)
}
