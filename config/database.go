package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

type Config struct {
	PostgresURL string
	RedisURL    string
	ServerPort  string
}

type Database struct {
	DB    *sql.DB
	Redis *redis.Client
}

func NewDatabase(config *Config) (*Database, error) {
	// Initialize PostgreSQL
	db, err := sql.Open("postgres", config.PostgresURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	// Initialize Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	return &Database{
		DB:    db,
		Redis: rdb,
	}, nil
}

func (d *Database) Close() {
	if err := d.DB.Close(); err != nil {
		log.Printf("Error closing PostgreSQL connection: %v", err)
	}
	if err := d.Redis.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	}
}
