package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-maps-api/model"
	"go-maps-api/repository"
)

type MapsHandler struct {
	Repository repository.Repository
}

func NewMapsHandler(repository repository.Repository) *MapsHandler {
	return &MapsHandler{
		Repository: repository,
	}
}

// GetRoutes
// @Summary Get routes between 2 points
// @Tags Maps
// @Produce  json
// @Param points query model.RouteRequest  true "Location latitude"
// @Router /maps/route [get]
func (h MapsHandler) GetRoutes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.RouteRequest
		if err := c.QueryParser(&params); err != nil {
			return c.Status(400).JSON(model.ErrorResponse{MESSAGE: err.Error()})
		}
		data, err := h.Repository.GetRoutes(params.StartLat, params.StartLng, params.EndLat, params.EndLng)
		if err != nil {
			return c.Status(400).JSON(model.ErrorResponse{MESSAGE: err.Error()})
		}
		return c.JSON(model.GetRouteResponse{
			MESSAGE: "get route success",
			DATA:    data,
		})
	}
}
