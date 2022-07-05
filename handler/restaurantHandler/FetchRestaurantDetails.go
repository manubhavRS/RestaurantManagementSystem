package restaurantHandler

import (
	"encoding/json"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/models/restaurantModels"
)

func FetchDishes(writer http.ResponseWriter, request *http.Request) {

	var restaurant restaurantModels.FetchRestaurantIDModel
	addErr := json.NewDecoder(request.Body).Decode(&restaurant)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	dishes, err := helper.FetchRestaurantDishes(restaurant.RestaurantID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(dishes)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(jsonData)
}
