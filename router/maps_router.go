package router

import (
	"github.com/gofiber/fiber/v2"
	"go-maps-api/handler"
	"go-maps-api/repository/openmaps"
)

func MapsRouter(router fiber.Router) {
	var mapsHandler = handler.NewMapsHandler(openmaps.New())

	group := router.Group("/maps")
	group.Get("/route", mapsHandler.GetRoutes())
}
