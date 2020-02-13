package app

import (
	"github.com/r-zareba/bookstore_users_api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
