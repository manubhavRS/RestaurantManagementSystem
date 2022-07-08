package userHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/userModels"
	"restaurantManagementSystem/utilities"
)

func AddPrivilege(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())
	log.Printf("Signed User: " + signedUser.Name)
	var user userModels.AddPrivilegeModel
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	roles, err := helper.FetchUserRoles(user.ID)
	if (user.Role == utilities.AdminRole && roles.Admin) || (user.Role == utilities.SubAdminRole && roles.SubAdmin) || (user.Role == utilities.UserRole && roles.User) {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = helper.AddPrivilege(user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
