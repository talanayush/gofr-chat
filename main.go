package main

import (
	"github.com/talanayush/gofr-chat/handler" // your health‐check (and other) handlers
	"gofr.dev/pkg/gofr"                       // the GoFr framework
)

func main() {
	// create a new GoFr application
	app := gofr.New()

	// wire up your health check endpoint
	app.GET("/health", handler.HealthCheck)

	// if you’ve generated a Swagger spec under ./docs/swagger.yaml,
	// GoFr can serve it for you.  By default it will render
	// ReDoc at /docs and the raw JSON/YAML at /docs/swagger.yaml
	app.GET("/swagger/*", gofr.SwaggerUIHandler)

	// start the server (defaults to :8000, or pass your own port: app.Run(":8080"))
	app.Run()
}
