package main

import (
	"github.com/Ayush09joshi/courseApp-ZopSmart.git/handler"
	"github.com/Ayush09joshi/courseApp-ZopSmart.git/store"
	"gofr.dev/pkg/gofr"
)

func main() {

	app := gofr.New()

	s := store.New()
	h := handler.New(s)

	// specifying the different routes supported by this service
	app.GET("/get", h.Get)
	app.POST("/create", h.Create)
	app.PUT("/update/{id}", h.Update)
	app.DELETE("/delete/{id}", h.Delete)

	app.Server.HTTP.Port = 3000
	app.Server.MetricsPort = 2113

	app.Start()

}
