package restaurantModels

import (
	"time"
)

type AllRestaurantModel struct {
	RestaurantID string    `db:"restaurant_id" json:"restaurantID"`
	Name         string    `db:"name" json:"name"`
	Latitude     string    `db:"latitude" json:"latitude"`
	Longitude    string    `db:"longitude" json:"longitude"`
	CreatedBy    string    `db:"created_by" json:"createdBy"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
}
