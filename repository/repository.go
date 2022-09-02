package repository

import "go-maps-api/model"

type Repository interface {
	GetRoutes(startLat float64, startLng float64, endLat float64, endLng float64) (*model.RouteData, error)
}
