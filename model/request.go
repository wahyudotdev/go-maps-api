package model

type RouteRequest struct {
	StartLat float64 `json:"start_lat" query:"start_lat"`
	StartLng float64 `json:"start_lng" query:"start_lng"`
	EndLat   float64 `json:"end_lat" query:"end_lat"`
	EndLng   float64 `json:"end_lng" query:"end_lng"`
}
