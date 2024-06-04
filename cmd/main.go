package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/marktlinn/RssScrape/db"
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

	db, err := db.NewDBConnecton(dbUrl)
	if err != nil {
		log.Fatalf("failed to connect to database instance: %s\n", err)
	}

	svr := server.NewServer(port, db)
	svr.RunServer()
}
