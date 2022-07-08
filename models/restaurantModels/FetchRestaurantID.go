package restaurantModels

type FetchRestaurantIDModel struct {
	RestaurantID string `db:"restaurant_id" json:"restaurantID"`
}
