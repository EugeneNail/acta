package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lib/pq"
	"log"
	"os"
)

const migrationsPath = "file:///migrations"
const commandUp = "up"
const commandDown = "down"
const noMigrationToRollbackError = "file does not exist"

// main creates the database when needed and runs the selected migration command.
func main() {
	applicationConfig, err := config.New()
	if err != nil {
		log.Fatal(fmt.Errorf("creating a config: %w", err))
	}

	if err := ensureDatabaseExists(applicationConfig.Postgres); err != nil {
		log.Fatal(fmt.Errorf("ensuring database exists: %w", err))
	}

	migrator, err := migrate.New(
		migrationsPath,
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
	defer migrator.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("creating a migrator: %w", err))
	}

	command := os.Getenv("MIGRATION_COMMAND")
	if command == "" {
		command = commandUp
	}

	switch command {
	case commandUp:
		if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(fmt.Errorf("running migrations up: %w", err))
		}
	case commandDown:
		err := migrator.Steps(-1)
		isActualError := err != nil && !errors.Is(err, migrate.ErrNoChange) && err.Error() != noMigrationToRollbackError
		if isActualError {
			log.Fatal(fmt.Errorf("running migrations down: %w", err))
		}
	default:
		log.Fatal(fmt.Errorf("reading migration command: unsupported command %q", command))
	}
}

// ensureDatabaseExists creates the target database when it is missing.
func ensureDatabaseExists(postgresConfig config.Postgres) error {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/postgres?sslmode=%s",
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.Host,
		postgresConfig.Port,
		postgresConfig.SslMode,
	))
	defer db.Close()
	if err != nil {
		return fmt.Errorf("opening postgres connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("pinging postgres connection: %w", err)
	}

	var exists bool
	if err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)",
		postgresConfig.Database,
	).Scan(&exists); err != nil {
		return fmt.Errorf("checking database existence: %w", err)
	}

	if exists {
		return nil
	}

	query := fmt.Sprintf("CREATE DATABASE %s", pq.QuoteIdentifier(postgresConfig.Database))
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("creating database: %w", err)
	}

	return nil
}
