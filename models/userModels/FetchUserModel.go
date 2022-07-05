package userModels

import (
	"restaurantManagementSystem/models"
	"time"
)

type FetchUserModel struct {
	ID        string                 `db:"user_id" json:"user_id"`
	Name      string                 `db:"name" json:"name"`
	Email     string                 `db:"email" json:"email"`
	Role      UserRoleModel          `db:"role" json:"role"`
	Location  []models.LocationModel `db:"location" json:"location"`
	CreatedAt time.Time              `db:"created_at" json:"created_at"`
}
