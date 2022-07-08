package restaurantModels

import (
	"github.com/volatiletech/null"
	"restaurantManagementSystem/models/DishModels"
	"time"
)

type RestaurantModel struct {
	RestaurantID string                 `db:"restaurant_id" json:"restaurantID"`
	Name         string                 `db:"name" json:"name"`
	Latitude     string                 `db:"latitude" json:"latitude"`
	Longitude    string                 `db:"longitude" json:"longitude"`
	Dishes       []DishModels.DishModel `db:"dishes" json:"dishes"`
	CreatedBy    string                 `db:"created_by" json:"createdBy"`
	CreatedAt    time.Time              `db:"created_at" json:"createdAt"`
	ArchivedAt   null.Time              `db:"archived_at" json:"archivedAt"`
}
