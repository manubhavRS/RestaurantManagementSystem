package restaurantHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/userModels"
)

func FetchSpecificRestaurant(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())
	log.Printf("Signed User: " + signedUser.Name)

	//if !signedUser.Role.SubAdmin && !signedUser.Role.Admin {
	//	writer.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	restaurants, userErr := helper.FetchSpecificRestaurant(signedUser.ID)
	if userErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(restaurants)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(jsonData)
	return
}
