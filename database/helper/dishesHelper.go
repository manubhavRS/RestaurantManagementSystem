package helper

import (
	"log"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/models/DishModels"
)

func FetchRestaurantDishes(restaurantID string) ([]DishModels.DishModel, error) {
	SQL := `SELECT dish_id, name, created_by 
			FROM dishes 
			WHERE restaurant_id=$1
			AND archived_at IS NULL;`

	dishes := make([]DishModels.DishModel, 0)
	log.Printf(restaurantID)
	err := database.Rms.Select(&dishes, SQL, restaurantID)
	if err != nil {
		return dishes, err
	}
	return dishes, nil
}
func FetchUserDishes(userID string) ([]DishModels.DishModel, error) {
	SQL := `SELECT dish_id, restaurant_id, name, created_by 
			FROM dishes 
			WHERE created_by=$1
			AND archived_at IS NULL;`

	dishes := make([]DishModels.DishModel, 0)
	err := database.Rms.Select(&dishes, SQL, userID)
	if err != nil {
		return dishes, err
	}

	return dishes, nil
}
func CreateDishes(userID, restaurantID string, dishes string) error {
	SQL := `INSERT INTO dishes(name, restaurant_id, created_by) 
			VALUES ($1,$2,$3) 
			RETURNING restaurant_id;`

	var dishID string
	err := database.Rms.QueryRow(SQL, dishes, restaurantID, userID).Scan(&dishID)
	if err != nil {
		return err
	}
	log.Printf("Dish: " + dishID + " has been added")

	return nil
}
