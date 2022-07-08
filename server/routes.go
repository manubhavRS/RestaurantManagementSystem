package server

import (
	"github.com/go-chi/chi"
	"net/http"
	"restaurantManagementSystem/handler"
	"restaurantManagementSystem/handler/dishHandler"
	"restaurantManagementSystem/handler/restaurantHandler"
	"restaurantManagementSystem/handler/userHandler"
	"restaurantManagementSystem/middlewareHandler"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()

	router.Route("/api", func(api chi.Router) {
		api.Post("/sign-in", handler.SigninUser)

		api.Route("/auth", func(auth chi.Router) {
			auth.Use(middlewareHandler.JWTAuthMiddleware)
			auth.Use(middlewareHandler.AccessControlMiddleware)
			auth.Route("/users", func(users chi.Router) {
				users.Post("/sign-up", userHandler.AddUser)
				users.Get("/all-users", userHandler.FetchUser)
				users.Get("/fetch-users", userHandler.FetchSpecificUser)
				users.Get("/fetch-subadmins", userHandler.FetchSubadminsUsers)
				users.Post("/fetch-distance", userHandler.UserDistance)
				users.Post("/add-roles", userHandler.AddPrivilege)
			})

			auth.Route("/restaurants", func(restaurants chi.Router) {
				restaurants.Post("/add-restaurant", restaurantHandler.AddRestaurant)
				restaurants.Get("/all-restaurants", restaurantHandler.FetchAllRestaurants)
				restaurants.Get("/fetch-restaurants", restaurantHandler.FetchSpecificRestaurant)
				restaurants.Post("/fetch-restaurant-dishes", restaurantHandler.FetchDishes)
			})
			auth.Route("/dishes", func(dishes chi.Router) {
				//dishes.Post("/fetch-restaurant-dishes", restaurantHandler.FetchDishes)
				dishes.Get("/fetch-user-dishes", dishHandler.FetchSpecificDishes)
				dishes.Post("/add-dish", dishHandler.AddDish)
				dishes.Post("/upload", dishHandler.UploadFile)
			})
		})
	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
