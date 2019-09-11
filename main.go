package main

import (
	"kiren-backend-go/app"

	"github.com/gin-gonic/gin"
)

func main() {
	log := app.InitLogger()

	if err := app.InitConfig(); err != nil {
		panic(err)
	}
	log.Infof("Initial config: %+v", app.CFG)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
