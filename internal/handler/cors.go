package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// corsMiddleware controls CORS in a request header.
func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}
