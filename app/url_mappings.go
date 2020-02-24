package app

import (
	"github.com/r-zareba/bookstore_users_api/controllers/ping"
	"github.com/r-zareba/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)   // Full update
	router.PATCH("/users/:user_id", users.Update) // Partial update
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
