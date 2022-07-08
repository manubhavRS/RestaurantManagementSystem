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
	var userLocationID models.UserLocationModel
	restaurantID := request.URL.Query().Get("restaurantID")
	if len(restaurantID) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&userLocationID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	distance.UserLocationID = userLocationID.UserLocationID
	distance.RestaurantID = restaurantID
	loc, err := FetchLocations(distance, signedUser)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	dist := utilities.DistanceCalculator(loc[0], loc[1], loc[2], loc[3])

	var ret = make(map[string]float64)
	ret["Distance_in_km"] = dist
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
func FetchLocations(distance models.DistanceModel, signedUser *userModels.UserModel) ([]float64, error) {
	var lat1, lng1, lat2, lng2 float64
	var dist []float64
	var err error
	if len(distance.UserLocationID) != 0 {
		location, err := helper.FetchLocation(distance.UserLocationID)
		if err != nil {
			return dist, err
		}

		lat1, err = strconv.ParseFloat(location.Latitude, 64)
		if err != nil {
			return dist, err
		}
		lng1, err = strconv.ParseFloat(location.Longitude, 64)
		if err != nil {
			return dist, err
		}
	} else {
		lat1, err = strconv.ParseFloat(signedUser.Location[0].Latitude, 64)
		if err != nil {
			return dist, err
		}

		lng1, err = strconv.ParseFloat(signedUser.Location[0].Longitude, 64)
		if err != nil {
			return dist, err
		}
	}

	location, err := helper.FetchRestaurantLocation(distance.RestaurantID)
	if err != nil {
		return dist, err
	}
	lat2, err = strconv.ParseFloat(location.Latitude, 64)
	if err != nil {
		return dist, err
	}
	lng2, err = strconv.ParseFloat(location.Longitude, 64)
	if err != nil {
		return dist, err
	}
	dist = append(dist, lat1)
	dist = append(dist, lng1)
	dist = append(dist, lat2)
	dist = append(dist, lng2)
	return dist, nil
}
