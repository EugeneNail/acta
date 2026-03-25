package main

import (
	"database/sql"
	"fmt"
	"github.com/EugeneNail/acta/journal/internal/infrastructure/config"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// main starts the journal HTTP server.
func main() {
	applicationConfig, err := config.New()
	if err != nil {
		log.Fatal(fmt.Errorf("creating a config: %w", err))
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s",
			applicationConfig.Postgres.User,
			applicationConfig.Postgres.Password,
			applicationConfig.Postgres.Host,
			applicationConfig.Postgres.Port,
			applicationConfig.Postgres.Database,
			applicationConfig.Postgres.SslMode,
		),
	)
	if err != nil {
		log.Fatal(fmt.Errorf("opening database connection: %w", err))
	}

	defer db.Close()

	server := http.NewServeMux()

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", applicationConfig.App.Port), server); err != nil {
		log.Fatal(fmt.Errorf("starting http server: %w", err))
	}
}
