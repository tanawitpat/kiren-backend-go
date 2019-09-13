package pingpong

import (
	"kiren-backend-go/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a logic of the /ping endpoint. It always returns pong.
func Ping(c *gin.Context) {
	logger := app.InitLogger()
	logger.Info("Ping endpoint called")
	c.JSON(http.StatusOK, "pong")
}
