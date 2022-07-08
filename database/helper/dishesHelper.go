package helper

import (
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
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
func FetchUserDishes(userID, restaurantID string) ([]DishModels.DishModel, error) {
	SQL := `SELECT dish_id, restaurant_id, name, created_by 
			FROM dishes 
			WHERE created_by=$1 AND restaurant_id=$2
			AND archived_at IS NULL;`

	dishes := make([]DishModels.DishModel, 0)
	err := database.Rms.Select(&dishes, SQL, userID, restaurantID)
	if err != nil {
		log.Printf("Error Fetching User Dishes")
		return dishes, err
	}
	return dishes, nil
}
func CreateDishes(userID, restaurantID string, dishes []string, tx *sqlx.Tx) error {
	psql := sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)
	insertBuilder := psql.Insert("dishes").Columns("name", "restaurant_id", "created_by")
	for _, dish := range dishes {
		insertBuilder.Values(dish, restaurantID, userID)
	}
	sql, args, err := insertBuilder.ToSql()
	if err != nil {
		log.Printf("Error Adding Dishes")
		return err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		log.Printf("Error Adding Dishes")
		return err
	}
	log.Printf("Dishes has been added")

	return nil
}
func CreateBulkDishes(dishList []DishModels.DishModel) error {
	psql := sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)
	insertBuilder := psql.Insert("dishes").Columns("name", "restaurant_id", "created_by")
	for _, dish := range dishList {
		insertBuilder.Values(dish.Name, dish.RestaurantID, dish.CreatedBy)
	}
	sql, args, err := insertBuilder.ToSql()
	if err != nil {
		log.Printf("Error Adding Dishes")
		return err
	}
	_, err = database.Rms.Exec(sql, args...)
	if err != nil {
		log.Printf("Error Adding Dishes")
		return err
	}
	log.Printf("Dishes has been added")

	return nil
}
