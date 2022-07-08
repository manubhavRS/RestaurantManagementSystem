package userModels

import "restaurantManagementSystem/models"

type AddUserModel struct {
	Name      string                 `db:"name" json:"name"`
	Email     string                 `db:"email" json:"email"`
	Password  string                 `db:"password" json:"password"`
	Role      []string               `db:"role" json:"role"`
	Location  []models.LocationModel `db:"location" json:"location"`
	CreatedBy string                 `db:"created_by" json:"createdBy"`
}
