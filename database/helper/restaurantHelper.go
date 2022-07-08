package helper

import (
	"github.com/jmoiron/sqlx"
	"log"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/models"
	"restaurantManagementSystem/models/restaurantModels"
)

func CreateRestaurant(restaurant restaurantModels.AddRestaurantModel, userID string, tx *sqlx.Tx) (string, error) {
	//language=SQL
	SQL := `INSERT INTO restaurant(name, latitude, longitude, created_by) 
			VALUES ($1,$2,$3,$4) 
			RETURNING restaurant_id;`
	var restaurantID string
	err := tx.Get(&restaurantID, SQL, restaurant.Name, restaurant.Latitude, restaurant.Longitude, userID)
	if err != nil {
		log.Printf("Error Adding Restaurant")
		return "", err
	}
	return restaurantID, nil
}

func AllRestaurants() ([]restaurantModels.AllRestaurantModel, error) {
	//language=SQL
	SQL := `SELECT restaurant_id, name, latitude, longitude, created_by, created_at 
			FROM restaurant 
			WHERE archived_at IS NULL;`
	restaurants := make([]restaurantModels.AllRestaurantModel, 0)
	errRow := database.Rms.Select(&restaurants, SQL)
	if errRow != nil {
		log.Printf("Error Retrieving Restaurants")
		return nil, errRow
	}
	return restaurants, nil
}

func FetchSpecificRestaurant(userID string) ([]restaurantModels.FetchRestaurantModel, error) {
	//language=SQL
	SQL := `SELECT restaurant_id, name, latitude,longitude, created_at 
			FROM restaurant 
			WHERE created_by=$1 
			AND archived_at IS NULL;`
	restaurants := make([]restaurantModels.FetchRestaurantModel, 0)
	err := database.Rms.Select(&restaurants, SQL, userID)
	if err != nil {
		log.Printf("Error Retrieving Restaurants")
		return restaurants, err
	}
	return restaurants, nil
}

func FetchRestaurantLocation(restaurantId string) (models.LocationModel, error) {
	//language=SQL
	SQL := `SELECT latitude,longitude
			FROM restaurant 
			WHERE restaurant_id=$1 
			AND archived_at IS NULL;`
	var location models.LocationModel
	err := database.Rms.Get(&location, SQL, restaurantId)
	if err != nil {
		log.Printf("Error Retrieving Restaurant location")
		return location, err
	}
	return location, nil
}
