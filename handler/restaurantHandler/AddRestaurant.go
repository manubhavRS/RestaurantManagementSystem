package restaurantHandler

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/restaurantModels"
	"restaurantManagementSystem/models/userModels"
)

func AddRestaurant(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())

	//if !signedUser.Role.Admin && !signedUser.Role.SubAdmin {
	//	writer.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	log.Printf("Signed User: " + signedUser.Name)
	var addRestaurant restaurantModels.AddRestaurantModel

	addErr := json.NewDecoder(request.Body).Decode(&addRestaurant)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {
		restaurantID, err := helper.CreateRestaurant(addRestaurant, signedUser.ID, tx)
		if err != nil {
			return err
		}
		err = helper.CreateDishes(signedUser.ID, restaurantID, addRestaurant.Dishes, tx)
		return err
	})

	if txErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(addRestaurant)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
