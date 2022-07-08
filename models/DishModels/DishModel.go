package DishModels

import (
	"github.com/volatiletech/null"
	"time"
)

type DishModel struct {
	DishID       string    `db:"dish_id" json:"dishID"`
	RestaurantID string    `db:"restaurant_id" json:"restaurantID"`
	Name         string    `db:"name" json:"name"`
	CreatedBy    string    `db:"created_by" json:"createdBy"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
	ArchivedAt   null.Time `db:"archived_at" json:"archivedAt"`
}
