package models

type DistanceModel struct {
	RestaurantID   string `db:"restaurant_id" json:"restaurant_id"`
	UserLocationID string `db:"location_id" json:"location_id"`
}
