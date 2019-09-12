package pingpong

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a logic of the /ping endpoint. It always returns pong.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
