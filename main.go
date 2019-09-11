package main

import (
	"kiren-backend-go/internal/app"
	"kiren-backend-go/internal/handler"
)

func main() {
	log := app.InitLogger()

	if err := app.InitConfig(); err != nil {
		panic(err)
	}
	log.Infof("Initial config: %+v", app.CFG)

	router := handler.NewRouter()
	router.Run()
}
