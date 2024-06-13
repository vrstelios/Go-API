package main

import (
	_ "context"
	"e/api"
	"e/config"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize PostgreSQL connection
	config.InitDB()
	defer config.CloseDB()

	// Retrieve the connection from the database package
	conn := config.GetDB()

	server := api.NewTourismHandler(conn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	if err := http.ListenAndServe(":"+port, server.Handler()); err != nil {
		log.Fatalf("Server error: %s - %v", err, err == http.ErrServerClosed)
	}
}
