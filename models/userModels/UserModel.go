package userModels

import (
	"github.com/volatiletech/null"
	"restaurantManagementSystem/models"
	"time"
)

type UserModel struct {
	ID         string                 `db:"user_id" json:"userID"`
	Name       string                 `db:"name" json:"name"`
	Email      string                 `db:"email" json:"email"`
	Password   string                 `db:"password" json:"password"`
	Role       UserRoleModel          `db:"role" json:"role"`
	Location   []models.LocationModel `db:"location" json:"location"`
	CreatedBy  string                 `db:"created_by" json:"createdBy"`
	CreatedAt  time.Time              `db:"created_at" json:"createdAt"`
	ArchivedAt null.Time              `db:"archived_at" json:"archivedAt"`
}
