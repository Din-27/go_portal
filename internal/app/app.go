package app

import (
	"fmt"
	"log"
	"net/http"
	"portal_service/internal/config"
	"portal_service/internal/transport/rest"
)

func Run() {
	fmt.Println("Run initialization whole app")
	db := config.DBinit()

	// Create an API handler which serves data from PlanetScale.
	handler := rest.NewHandler(db)

	// Start an HTTP API server.
	const addr = ":8080"
	log.Printf("successfully connected to PlanetScale, starting HTTP server on %q", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
