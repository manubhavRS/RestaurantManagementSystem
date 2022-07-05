package restaurantModels

import (
	"restaurantManagementSystem/models/DishModels"
	"time"
)

type FetchRestaurantModel struct {
	RestaurantID string                 `db:"restaurant_id" json:"restaurant_id"`
	Name         string                 `db:"name" json:"name"`
	Latitude     string                 `db:"latitude" json:"latitude"`
	Longitude    string                 `db:"longitude" json:"longitude"`
	Dishes       []DishModels.DishModel `db:"dishes" json:"dishes"`
	CreatedBy    string                 `db:"created_by" json:"created_by"`
	CreatedAt    time.Time              `db:"created_at" json:"created_at"`
}
