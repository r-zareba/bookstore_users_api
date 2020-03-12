package app

import (
	"github.com/gin-gonic/gin"
	"github.com/r-zareba/bookstore_users_api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Starting application...")
	router.Run(":8081")

}
