package database

import (
	"context"
	"e/model"
	pgx "github.com/jackc/pgx/v4"
)

func PostTourism(conn *pgx.Conn, tourism model.Tourism) (model.Tourism, error) {
	_, err := conn.Exec(context.Background(), "INSERT INTO tourism (id, tourism_gr, tourism, area_name) VALUES ($1, $2, $3, $4)",
		tourism.Id, tourism.TourismGR, tourism.Tourism, tourism.AreaName)
	if err != nil {
		return tourism, err
	}

	return tourism, nil
}

func GetTourism(conn *pgx.Conn, id string) (model.Tourism, error) {
	var tourism model.Tourism
	err := conn.QueryRow(context.Background(), "SELECT id, tourism_gr, tourism, area_name FROM tourism WHERE id=$1", id).Scan(&tourism.Id, &tourism.TourismGR, &tourism.Tourism, &tourism.AreaName)
	if err != nil {
		return tourism, err
	}

	return tourism, nil
}

func PutTourism(conn *pgx.Conn, id string, tourism model.Tourism) error {
	_, err := conn.Exec(context.Background(), "UPDATE tourism SET tourism_gr=$1, tourism=$2, area_name=$3 WHERE id=$4",
		tourism.TourismGR, tourism.Tourism, tourism.AreaName, id)
	return err
}

func DelTourism(conn *pgx.Conn, id string) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM tourism WHERE id=$1", id)
	return err
}
