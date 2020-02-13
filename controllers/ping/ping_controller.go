package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(context *gin.Context) {
	context.String(http.StatusOK, "PONG")
}
