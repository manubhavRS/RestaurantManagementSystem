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

	if signedUser.Role.SubAdmin || signedUser.Role.Admin {
		dishes, err := helper.FetchUserDishes(signedUser.ID)
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
	writer.WriteHeader(http.StatusUnauthorized)
	return
}
