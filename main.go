package main

import (
	"github.com/talanayush/gofr-chat/handler" 
	"gofr.dev/pkg/gofr"                       
)

func main() {

	app := gofr.New()


	app.GET("/health", handler.HealthCheck)

	app.GET("/swagger/*", gofr.SwaggerUIHandler)


	app.Run()
}
