package models

type DistanceModel struct {
	RestaurantID   string `db:"restaurant_id" json:"restaurantID"`
	UserLocationID string `db:"location_id" json:"locationID"`
}
type UserLocationModel struct {
	UserLocationID string `db:"location_id" json:"locationID"`
}
