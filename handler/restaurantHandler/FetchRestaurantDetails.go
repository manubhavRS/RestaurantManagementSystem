package restaurantHandler

import (
	"encoding/json"
	"net/http"
	"restaurantManagementSystem/database/helper"
)

func FetchDishes(writer http.ResponseWriter, request *http.Request) {

	restaurantID := request.URL.Query().Get("restaurantID")
	if len(restaurantID) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	dishes, err := helper.FetchRestaurantDishes(restaurantID)
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
