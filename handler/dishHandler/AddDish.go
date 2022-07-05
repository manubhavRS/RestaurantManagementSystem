package dishHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/DishModels"
	"restaurantManagementSystem/models/userModels"
)

func AddDish(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())

	if !signedUser.Role.Admin && !signedUser.Role.SubAdmin {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Printf("Signed User: " + signedUser.Name)
	var addDish DishModels.DishModel
	addErr := json.NewDecoder(request.Body).Decode(&addDish)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := helper.CreateDishes(signedUser.ID, addDish.RestaurantID, addDish.Name)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(addDish)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
