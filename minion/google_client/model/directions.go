package model

type GoogleClientGetDirectionsRequest struct {
	Origin      *string `json:"origin,omitempty" url:"origin"`
	Destination *string `json:"destination,omitempty" url:"destination"`
	Key         *string `json:"key,omitempty" url:"key"`
}

type GoogleClientGetDirectionsResponse struct {
	Status            *string                 `json:"status"`
	GeocodedWaypoints *[]GeocodedWaypointData `json:"geocoded_waypoints"`
	Routes            *[]RouteData            `json:"routes"`
}

type GeocodedWaypointData struct {
	GeocoderStatus *string   `json:"geocoder_status"`
	PlaceID        *string   `json:"place_id"`
	Types          *[]string `json:"types"`
}

type RouteData struct {
	Summary *string    `json:"summary"`
	Legs    *[]LegData `json:"legs"`
}

type LegData struct {
	Steps         *[]StepData   `json:"steps"`
	Duration      *DurationData `json:"duration"`
	Distance      *DistanceData `json:"distance"`
	StartLocation *LocationData `json:"start_location"`
	EndLocation   *LocationData `json:"end_location"`
}

type StepData struct {
	Duration      *DurationData `json:"duration"`
	Distance      *DistanceData `json:"distance"`
	StartLocation *LocationData `json:"start_location"`
	EndLocation   *LocationData `json:"end_location"`
}

type DurationData struct {
	Value *int    `json:"value"`
	Text  *string `json:"text"`
}

type DistanceData struct {
	Value *int    `json:"value"`
	Text  *string `json:"text"`
}

type LocationData struct {
	Latitude  *float64 `json:"lat"`
	Longitude *float64 `json:"lng"`
}
