package userHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/userModels"
)

func AddUser(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())
	log.Printf("Signed User: " + signedUser.Name)
	log.Printf("Hello")
	var addUser userModels.AddUserModel
	addErr := json.NewDecoder(request.Body).Decode(&addUser)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, role := range addUser.Role {
		if (role == "admin" && !signedUser.Role.Admin) || (role == "sub-admin" && !signedUser.Role.Admin) || (role == "user" && (!signedUser.Role.Admin && !signedUser.Role.SubAdmin)) {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	addUser.CreatedBy = signedUser.ID
	log.Printf(addUser.Name)
	userID, err := helper.CreateUser(addUser)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("User: " + userID + " has been added")

	err = helper.AddLocation(addUser, userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Locations has been added")

	err = helper.AddUserRole(addUser, userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print("User roles have been added")

	jsonData, jsonErr := json.Marshal(addUser)

	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Write(jsonData)
}
