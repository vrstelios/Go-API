package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

var conn *pgx.Conn

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "postgres"
)

func InitDB() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	conn, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

func GetDB() *pgx.Conn {
	return conn
}

func CloseDB() {
	conn.Close(context.Background())
}
