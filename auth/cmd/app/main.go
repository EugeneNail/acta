package main

import (
	"database/sql"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/application/create_user"
	"github.com/EugeneNail/acta/auth/internal/application/login_user"
	"github.com/EugeneNail/acta/auth/internal/application/refresh_access_token"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/config"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/repository/postgres"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/token"
	transportHttp "github.com/EugeneNail/acta/auth/internal/transport/http"
	"github.com/EugeneNail/acta/auth/internal/transport/http/middleware"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// main starts the auth HTTP server.
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

	userRepository := postgres.NewUserRepository(db)
	createUserHandler := create_user.NewHandler(userRepository)
	tokenProvider := token.NewProvider()
	loginUserHandler := login_user.NewHandler(userRepository, tokenProvider)
	refreshAccessTokenHandler := refresh_access_token.NewHandler(userRepository, tokenProvider)

	server := http.NewServeMux()
	httpHandler := transportHttp.NewHandler(createUserHandler, loginUserHandler, refreshAccessTokenHandler)

	server.HandleFunc("GET  /api/v1/auth", httpHandler.Ping)
	server.HandleFunc("POST /api/v1/auth/signup", middleware.WriteJsonResponse(httpHandler.Signup))
	server.HandleFunc("POST /api/v1/auth/login", middleware.WriteJsonResponse(httpHandler.Login))
	server.HandleFunc("POST /api/v1/auth/refresh", middleware.WriteJsonResponse(httpHandler.Refresh))

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", applicationConfig.App.Port), server); err != nil {
		log.Fatal(fmt.Errorf("starting http server: %w", err))
	}
}
