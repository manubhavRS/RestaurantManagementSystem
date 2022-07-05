package restaurantModels

type AddRestaurantModel struct {
	Name      string   `db:"name" json:"name"`
	Latitude  string   `db:"latitude" json:"latitude"`
	Longitude string   `db:"longitude" json:"longitude"`
	Dishes    []string `db:"dishes" json:"dishes"`
}
