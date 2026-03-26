package main

import (
	"database/sql"
	"fmt"
	"github.com/EugeneNail/acta/journal/internal/application/create_habit"
	"github.com/EugeneNail/acta/journal/internal/application/delete_habit"
	"github.com/EugeneNail/acta/journal/internal/application/update_habit"
	"github.com/EugeneNail/acta/journal/internal/infrastructure/config"
	"github.com/EugeneNail/acta/journal/internal/infrastructure/repository/postgres"
	transportHttp "github.com/EugeneNail/acta/journal/internal/transport/http"
	"github.com/EugeneNail/acta/journal/internal/transport/http/middleware"
	libHttpMiddleware "github.com/EugeneNail/acta/lib-common/pkg/http/middleware"
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

	habitRepository := postgres.NewHabitRepository(db)

	createHabitHandler := create_habit.NewHandler(habitRepository)
	deleteHabitHandler := delete_habit.NewHandler(habitRepository)
	updateHabitHandler := update_habit.NewHandler(habitRepository)

	server := http.NewServeMux()
	httpHandler := transportHttp.NewHandler(
		createHabitHandler,
		deleteHabitHandler,
		updateHabitHandler,
	)

	server.HandleFunc("POST /api/v1/journal/habits", middleware.Authenticate(libHttpMiddleware.WriteJsonResponse(httpHandler.CreateHabit)))
	server.HandleFunc("DELETE /api/v1/journal/habits/{uuid}", middleware.Authenticate(libHttpMiddleware.WriteJsonResponse(httpHandler.DeleteHabit)))
	server.HandleFunc("PUT /api/v1/journal/habits/{uuid}", middleware.Authenticate(libHttpMiddleware.WriteJsonResponse(httpHandler.UpdateHabit)))

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", applicationConfig.App.Port), server); err != nil {
		log.Fatal(fmt.Errorf("starting http server: %w", err))
	}
}
