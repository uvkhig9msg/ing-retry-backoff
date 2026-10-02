// Command ing-retry-backoff serves the HTTP API described in README.md.
package main

import (
	"log"
	"os"

	"github.com/uvkhig9msg/ing-retry-backoff/internal/api"
	"github.com/uvkhig9msg/ing-retry-backoff/internal/store"
)

func main() {
	address := os.Getenv("ADDR")
	if address == "" {
		address = "127.0.0.1:8080"
	}
	databasePath := os.Getenv("DB_PATH")
	if databasePath == "" {
		databasePath = "ing-retry-backoff.db"
	}

	st, err := store.Open(databasePath)
	if err != nil {
		log.Fatalf("open store: %v", err)
	}
	defer st.Close()

	if err := api.NewRouter(st).Run(address); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
