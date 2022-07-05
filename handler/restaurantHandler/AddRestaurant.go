package restaurantHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/restaurantModels"
	"restaurantManagementSystem/models/userModels"
)

func AddRestaurant(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())

	if !signedUser.Role.Admin && !signedUser.Role.SubAdmin {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Printf("Signed User: " + signedUser.Name)
	var addRestaurant restaurantModels.AddRestaurantModel

	addErr := json.NewDecoder(request.Body).Decode(&addRestaurant)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	restaurantID, err := helper.CreateRestaurant(addRestaurant, signedUser.ID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Restaurant: " + restaurantID + " has been added")

	jsonData, jsonErr := json.Marshal(addRestaurant)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
