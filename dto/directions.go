package dto

type GetDirectionsRequest struct {
	RequestID     string      `json:"request-id"`
	StartLat string `json:"start-lat"`
	StartLon   string `json:"start-lon"`
	EndLat    string `json:"end-lat"`
	EndLon  string `json:"end-lon"`
}

type GetDirectionsResponse struct {
	RequestID  string        `json:"request-id"`
	Directions []Coordinates `json:"directions"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
