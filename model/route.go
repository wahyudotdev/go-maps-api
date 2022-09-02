package model

type RouteResponse struct {
	Code string   `json:"code"`
	Data []Routes `json:"routes"`
}

type Routes struct {
	Legs []Legs `json:"legs"`
}

type Legs struct {
	Steps    []Step  `json:"steps"`
	Summary  string  `json:"summary"`
	Duration float64 `json:"duration"`
	Distance float64 `json:"distance"`
}

type Step struct {
	Maneuver Maneuver `json:"maneuver"`
}

type Maneuver struct {
	BearingAfter  int64     `json:"bearing_after"`
	BearingBefore int64     `json:"bearing_before"`
	Modifier      string    `json:"modifier"`
	Location      []float64 `json:"location"`
}

type GetRouteResponse struct {
	MESSAGE string     `json:"MESSAGE"`
	DATA    *RouteData `json:"DATA"`
}

type RouteData struct {
	Duration float64      `json:"duration"`
	Distance float64      `json:"distance"`
	Points   []RoutePoint `json:"points"`
}

type RoutePoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
