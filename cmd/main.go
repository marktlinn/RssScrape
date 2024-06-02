package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/marktlinn/RssScrape/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("failed to find port: PORT env variable must be set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalln("failed to connect to DB: DB env variable must be set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("failed to connect to Postgres instance: %s\n", err)
	}

	log.Printf("connected to postgres: %v\n", conn)

	server.Server(port)
}
