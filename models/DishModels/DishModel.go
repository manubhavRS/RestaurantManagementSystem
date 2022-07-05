package DishModels

import (
	"github.com/volatiletech/null"
	"time"
)

type DishModel struct {
	DishID       string    `db:"dish_id" json:"dish_id"`
	RestaurantID string    `db:"restaurant_id" json:"restaurant_id"`
	Name         string    `db:"name" json:"name"`
	CreatedBy    string    `db:"created_by" json:"created_by"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	ArchivedAt   null.Time `db:"archived_at" json:"archived_at"`
}
