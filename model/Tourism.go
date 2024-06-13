package model

import "github.com/google/uuid"

type Tourism struct {
	Id        uuid.UUID `json:"id"`
	TourismGR float64   `json:"tourism_gr"`
	Tourism   float64   `json:"tourism"`
	AreaName  string    `json:"area_name"`
}
