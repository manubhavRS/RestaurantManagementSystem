package models

type LocationModel struct {
	LocationID string `db:"location_id" json:"location_id"`
	Latitude   string `db:"latitude" json:"latitude"`
	Longitude  string `db:"longitude" json:"longitude"`
}
