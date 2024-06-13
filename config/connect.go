package config

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

var conn *pgx.Conn

func InitDB() {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/postgres")
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
