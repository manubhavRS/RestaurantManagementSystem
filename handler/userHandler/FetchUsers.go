package userHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/userModels"
)

func FetchUser(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())
	log.Printf("Signed User: " + signedUser.Name)
	//if !signedUser.Role.Admin {
	//	writer.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	users, userErr := helper.FetchAllUser(signedUser.ID)
	if userErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(users)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(jsonData)
	return
}
