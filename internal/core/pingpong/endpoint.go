package pingpong

import (
	"kiren-backend-go/internal/app"
	"net/http"
)

// Ping is a logic of the /ping endpoint. It always returns pong.
func Ping() (string, int) {
	logger := app.InitLogger()
	logger.Info("Ping endpoint called")
	return "pong", http.StatusOK
}
