package utilities

const Secretkey string = "SuperSecretKey"
const ContextUserKey string = "user"
const ContextRefreshToken string = "refreshToken"
const AdminRole = "admin"
const SubAdminRole = "sub-admin"
const UserRole = "user"

const Path1 = "/api/auth/users/sign-up"
const Path2 = "/api/auth/users/all-users" //admin
const Path3 = "/api/auth/users/fetch-users"
const Path4 = "/api/auth/users/fetch-subadmins" //admin
const Path5 = "/api/auth/users/fetch-distance"
const Path6 = "/api/auth/users/add-roles" //admin
const Path7 = "/api/auth/restaurants/add-restaurant"
const Path8 = "/api/auth/restaurants/all-restaurants" //admin
const Path9 = "/api/auth/restaurants/fetch-restaurants"
const Path10 = "/api/auth/restaurants/fetch-restaurant-dishes"
const Path11 = "/api/auth/dishes/fetch-user-dishes"
const Path12 = "/api/auth/dishes/fetch-restaurant-dishes"
const Path13 = "/api/auth/dishes/add-dish"
const Path14 = "/api/auth/dishes/upload"

//auth.Route("/dishes", func(dishes chi.Router) {
//	//dishes.Post("/fetch-restaurant-dishes", restaurantHandler.FetchDishes)
//	dishes.Get("/fetch-user-dishes", dishHandler.FetchSpecificDishes)
//	dishes.Post("/add-dish", dishHandler.AddDish)
//	dishes.Post("/upload", dishHandler.UploadFile)
