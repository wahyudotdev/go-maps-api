package openmaps

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-maps-api/model"
)

type OpenMaps struct {
	client *resty.Client
}

func New() OpenMaps {
	client := resty.New()
	return OpenMaps{
		client: client,
	}
}

func (r OpenMaps) GetRoutes(startLat float64, startLng float64, endLat float64, endLng float64) (*model.RouteData, error) {
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=false&steps=true", startLng, startLat, endLng, endLat)
	resp, err := r.client.R().
		EnableTrace().
		Get(url)
	if err != nil {
		return nil, err
	}
	var data model.RouteResponse
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, err
	}
	result := make([]model.RoutePoint, 0)
	for _, d := range data.Data[0].Legs[0].Steps {
		item := model.RoutePoint{
			Lat: d.Maneuver.Location[0],
			Lng: d.Maneuver.Location[1],
		}
		result = append(result, item)
	}
	return &model.RouteData{
		Distance: data.Data[0].Legs[0].Distance,
		Duration: data.Data[0].Legs[0].Duration,
		Points:   result,
	}, nil
}
