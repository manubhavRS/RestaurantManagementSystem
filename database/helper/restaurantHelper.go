package helper

import (
	"log"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/models"
	"restaurantManagementSystem/models/restaurantModels"
)

func CreateRestaurant(restaurant restaurantModels.AddRestaurantModel, userID string) (string, error) {
	SQL := `INSERT INTO restaurant(name, latitude, longitude, created_by) 
			VALUES ($1,$2,$3,$4) 
			RETURNING restaurant_id;`
	var restaurantID string
	err := database.Rms.QueryRow(SQL, restaurant.Name, restaurant.Latitude, restaurant.Longitude, userID).Scan(&restaurantID)
	if err != nil {
		return "", err
	}

	log.Printf("Resaturant:" + restaurantID + "has been added")

	for _, dish := range restaurant.Dishes {
		err = CreateDishes(userID, restaurantID, dish)
		if err != nil {
			return "", err
		}
	}

	return restaurantID, nil
}

func AllRestaurants() ([]restaurantModels.AllRestaurantModel, error) {
	SQL := `SELECT restaurant_id, name, latitude, longitude, created_by, created_at 
			FROM restaurant 
			WHERE archived_at IS NULL;`
	restaurants := make([]restaurantModels.AllRestaurantModel, 0)

	rows, errRow := database.Rms.Queryx(SQL)
	if errRow != nil {
		return nil, errRow
	}
	for rows.Next() {
		var u restaurantModels.AllRestaurantModel
		rows.Scan(&u.RestaurantID, &u.Name, &u.Latitude, &u.Longitude, &u.CreatedBy, &u.CreatedAt)

		dishes, err := FetchRestaurantDishes(u.RestaurantID)
		if err != nil {
			return nil, err
		}
		u.Dishes = dishes

		restaurants = append(restaurants, u)
	}

	return restaurants, nil
}

func FetchSpecificRestaurant(userID string) ([]restaurantModels.FetchRestaurantModel, error) {
	SQL := `SELECT restaurant_id, name, latitude,longitude, created_at 
			FROM restaurant 
			WHERE created_by=$1 
			AND archived_at IS NULL;`

	restaurants := make([]restaurantModels.FetchRestaurantModel, 0)
	err := database.Rms.Select(&restaurants, SQL, userID)
	if err != nil {
		return restaurants, err
	}
	for i, restaurant := range restaurants {
		dishes, err := FetchRestaurantDishes(restaurant.RestaurantID)
		if err != nil {
			return nil, err
		}
		restaurants[i].Dishes = dishes
	}
	return restaurants, nil
}

func FetchRestaurantLocation(restaurantId string) (models.LocationModel, error) {
	SQL := `SELECT latitude,longitude
			FROM restaurant 
			WHERE restaurant_id=$1 
			AND archived_at IS NULL;`

	var location models.LocationModel
	err := database.Rms.QueryRow(SQL, restaurantId).Scan(&location.Latitude, &location.Longitude)
	if err != nil {
		return location, err
	}

	return location, nil
}
