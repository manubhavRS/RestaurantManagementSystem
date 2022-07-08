package utilities

import (
	"log"
	"restaurantManagementSystem/models/DishModels"
)

func CreateDishList(data [][]string) []DishModels.DishModel {
	var dishList []DishModels.DishModel
	var dish DishModels.DishModel
	for i, val := range data {
		if i > 1 {
			dish.RestaurantID = val[0]
			dish.Name = val[1]
			dish.CreatedBy = val[2]
			log.Printf(dish.RestaurantID + dish.Name + dish.CreatedBy)
			dishList = append(dishList, dish)
		}
	}
	return dishList
}
