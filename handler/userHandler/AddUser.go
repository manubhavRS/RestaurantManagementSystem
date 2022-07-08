package userHandler

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/userModels"
)

func AddUser(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())
	log.Printf("Signed User: " + signedUser.Name)
	var addUser userModels.AddUserModel
	addErr := json.NewDecoder(request.Body).Decode(&addUser)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	//
	//for _, role := range addUser.Role {
	//	if role == utilities.AdminRole || role == utilities.SubAdminRole {
	//		writer.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	addUser.CreatedBy = signedUser.ID
	txErr := database.Tx(func(tx *sqlx.Tx) error {
		userID, err := helper.CreateUser(addUser, tx)
		err = helper.AddLocation(addUser, userID, tx)
		err = helper.AddUserRole(addUser, userID, tx)
		return err
	})
	if txErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(addUser)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
