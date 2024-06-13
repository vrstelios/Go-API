package database

import (
	"context"
	"e/model"
	pgx "github.com/jackc/pgx/v4"
)

func PostTourism(conn *pgx.Conn, t model.Tourism) (model.Tourism, error) {
	row := conn.QueryRow(context.Background(), "INSERT INTO tourism (id, tourism_gr, tourism, area_name) VALUES ($1, $2, $3, $4)",
		t.Id, t.TourismGR, t.Tourism, t.AreaName)
	err := row.Scan(&t.Id, &t.TourismGR, &t.Tourism, &t.AreaName)
	if err != nil {
		return t, err
	}
	return t, nil

}

func GetTourism(conn *pgx.Conn, id string) (model.Tourism, error) {
	var t model.Tourism
	row := conn.QueryRow(context.Background(), "SELECT id, tourism_gr, tourism, area_name FROM tourism WHERE id=$1", id)
	err := row.Scan(&t.Id, &t.TourismGR, &t.Tourism, &t.AreaName)
	if err != nil {
		return t, err
	}
	return t, nil
}

func PutTourism(conn *pgx.Conn, id string, t model.Tourism) error {
	_, err := conn.Exec(context.Background(), "UPDATE tourism SET tourism_gr=$1, tourism=$2, area_name=$3 WHERE id=$4",
		t.TourismGR, t.Tourism, t.AreaName, id)
	return err
}

func DelTourism(conn *pgx.Conn, id string) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM tourism WHERE id=$1", id)
	return err
}
