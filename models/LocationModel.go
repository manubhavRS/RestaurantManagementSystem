package models

type LocationModel struct {
	LocationID string `db:"location_id" json:"locationID"`
	Latitude   string `db:"latitude" json:"latitude"`
	Longitude  string `db:"longitude" json:"longitude"`
}
