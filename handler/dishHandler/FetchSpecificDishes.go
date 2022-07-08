package dishHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/userModels"
)

func FetchSpecificDishes(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())
	log.Printf("Signed User: " + signedUser.Name)
	//if !signedUser.Role.SubAdmin && !signedUser.Role.Admin {
	//	writer.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	restaurantID := request.URL.Query().Get("restaurantID")
	if len(restaurantID) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	dishes, err := helper.FetchUserDishes(signedUser.ID, restaurantID)
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
	return
}
