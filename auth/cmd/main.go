package main

import (
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/config"
	http2 "github.com/EugeneNail/acta/auth/internal/transport/http"
	"log"
	"net/http"
)

func main() {
	applicationConfig, err := config.New()
	if err != nil {
		log.Fatal(fmt.Errorf("creating a config: %w", err))
	}

	server := http.NewServeMux()
	httpHandler := http2.NewHandler()

	server.HandleFunc("GET  /", httpHandler.Ping)

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", applicationConfig.App.Port), server); err != nil {
		log.Fatal(err)
	}
}
