package dishHandler

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/DishModels"
	"restaurantManagementSystem/models/userModels"
)

func AddDish(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())

	//if !signedUser.Role.Admin && !signedUser.Role.SubAdmin {
	//	writer.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	log.Printf("Signed User: " + signedUser.Name)
	restaurantID := request.URL.Query().Get("restaurantID")
	if len(restaurantID) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var addDish DishModels.DishModel
	addErr := json.NewDecoder(request.Body).Decode(&addDish)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	addDish.RestaurantID = restaurantID
	var dishes []string
	dishes = append(dishes, addDish.Name)
	txErr := database.Tx(func(tx *sqlx.Tx) error {
		err := helper.CreateDishes(signedUser.ID, addDish.RestaurantID, dishes, tx)
		return err
	})
	if txErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(addDish)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
