package app

import (
	"github.com/r-zareba/bookstore_users_api/controllers/ping"
	"github.com/r-zareba/bookstore_users_api/controllers/users_controller"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users_controller.Get)
	router.POST("/users", users_controller.Create)
	router.PUT("/users/:user_id", users_controller.Update)   // Full update
	router.PATCH("/users/:user_id", users_controller.Update) // Partial update
	router.DELETE("/users/:user_id", users_controller.Delete)
	router.GET("/internal/users/search", users_controller.Search)
}
