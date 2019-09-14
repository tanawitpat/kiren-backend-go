package main

import (
	"kiren-backend-go/internal/app"
	"kiren-backend-go/internal/handler"
	"strconv"
)

func main() {
	log := app.InitLogger()

	// Read service configs from a config file
	if err := app.InitConfig(); err != nil {
		panic(err)
	}
	log.Infof("Initial app config: %+v", app.CFG)
	log.Infof("Initial error config: %+v", app.ERR)

	router := handler.NewRouter()
	router.Run(":" + strconv.Itoa(app.CFG.App.Port))
}
