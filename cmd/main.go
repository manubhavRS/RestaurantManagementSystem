package main

import (
	"fmt"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/server"
)

func main() {
	err := database.ConnectAndMigrate("localhost", "5434", "restaurant", "local", "local", database.SSLModeDisable)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	srv := server.SetupRoutes()
	err = srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
