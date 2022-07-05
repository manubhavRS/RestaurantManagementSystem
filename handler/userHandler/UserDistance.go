package userHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/middlewareHandler"
	"restaurantManagementSystem/models"
	"restaurantManagementSystem/models/userModels"
	"restaurantManagementSystem/utilities"
	"strconv"
)

func UserDistance(writer http.ResponseWriter, request *http.Request) {
	var signedUser *userModels.UserModel
	signedUser = middlewareHandler.UserFromContext(request.Context())
	log.Printf("Signed User: " + signedUser.Name)

	var distance models.DistanceModel
	err := json.NewDecoder(request.Body).Decode(&distance)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var lat1, lng1, lat2, lng2 float64
	if len(distance.UserLocationID) != 0 {
		location, err := helper.FetchLocation(distance.UserLocationID)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		lat1, err = strconv.ParseFloat(location.Latitude, 64)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		lng1, err = strconv.ParseFloat(location.Longitude, 64)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		lat1, err = strconv.ParseFloat(signedUser.Location[0].Latitude, 64)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		lng1, err = strconv.ParseFloat(signedUser.Location[0].Longitude, 64)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	location, err := helper.FetchRestaurantLocation(distance.RestaurantID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	lat2, err = strconv.ParseFloat(location.Latitude, 64)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	lng2, err = strconv.ParseFloat(location.Longitude, 64)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	dist := utilities.DistanceCalculator(lat1, lat2, lng1, lng2)

	var ret = make(map[string]float64)
	ret["Distance in KM"] = dist
	jsonString, jsonErr := json.Marshal(ret)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = writer.Write(jsonString)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}
