package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/marktlinn/RssScrape/internal/database"
)

type Config struct {
	DB *database.Queries
}

func NewDBConnecton(dbUrl string) (*Config, error) {
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres instance: %w", err)
	}

	log.Printf("connected to postgres: %v\n", conn)

	queries := database.New(conn)
	if queries == nil {
		return nil, fmt.Errorf("failed to connect to database and retrieve query object")
	}

	return &Config{DB: queries}, nil
}
