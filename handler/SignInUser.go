package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models/userModels"
	"restaurantManagementSystem/utilities"
	"time"
)

func SigninUser(w http.ResponseWriter, r *http.Request) {

	var creds userModels.SignInUserModel
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := helper.SignInCredential(creds.Email)
	ok := utilities.CheckPasswordHash(creds.Password, user.Password)
	if !ok {
		//log.Printf("Password Incorrect")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString, err := middlewareHandler.GenerateJWT(user)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["Name"] = "token"
	resp["Value"] = tokenString
	resp["Expires"] = (time.Now().Add(5 * time.Hour)).String()
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResponse)
}
