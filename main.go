package main

import (
	_ "context"
	"e/api"
	"e/config"
	_ "e/docs"
	"log"
	"net/http"
	"os"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8000
// @BasePath /api
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
	if err := http.ListenAndServe(":"+port, server.Router()); err != nil {
		log.Fatalf("Server error: %s - %v", err, err == http.ErrServerClosed)
	}
}
