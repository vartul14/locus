package logic

import (
	"vartul14/locus/minion/google_client/model"
	"fmt"
	"vartul14/locus/dto"
	"vartul14/locus/minion"
	"vartul14/locus/common"
	"math"
)

func GetDirections(request dto.GetDirectionsRequest) (*dto.GetDirectionsResponse, error) {
	//Get Direction from Google Directions API
	response, errGetDirections := getDirections(request)
	if errGetDirections != nil {
		fmt.Println("Error getting directions from the external API | Request ID = %v | Error = %v", request.RequestID, errGetDirections.Error())
		return nil, errGetDirections
	}
	fmt.Printf("Response from Directions API %v", response)

	//Divide the directions into equally spaced coordinates
	equallySpacedCoordinates, errGetEquallySpacedCoordinates := getEquallySpacedCoordinates(request.RequestID, response)
	if errGetEquallySpacedCoordinates != nil {
		fmt.Println("Error getting equally spaced coordinates | RequestID = %v | Error = %v", request.RequestID, errGetEquallySpacedCoordinates.Error())
		return nil, errGetEquallySpacedCoordinates
	}

	getDirectionsResponse := &dto.GetDirectionsResponse{
		RequestID: request.RequestID,
		Directions: equallySpacedCoordinates,
	}

	fmt.Printf("Successfully fetched coordiantes | RequestID = %v", request.RequestID)
	return getDirectionsResponse, nil
}

func getDirections(request dto.GetDirectionsRequest) (model.GoogleClientGetDirectionsResponse, error) {
	getDirectionsRequest := createGetDirectionsRequest(request)

	response, _, err := minion.GetGoogleClient().GetDirections(request.RequestID, getDirectionsRequest)
	return response, err
}

func createGetDirectionsRequest(request dto.GetDirectionsRequest) map[string]string {
	origin := request.StartLat + "," + request.StartLon
	destination := request.EndLat + "," + request.EndLon

	// getDirectionsRequest := model.GoogleClientGetDirectionsRequest{
	// 	Origin:      origin,
	// 	Destination: destination,
	// 	Key:         "fgfgdfg",
	// }

	getDirectionsRequest := map[string]string{
		"origin":      origin,
		"destination": destination,
		"key":         "AIzaSyAb8ohmBXqtK4y2_a5CFnFnfLGiOsuwjIo",
	}

	return getDirectionsRequest
}

func getEquallySpacedCoordinates(requestID string, req model.GoogleClientGetDirectionsResponse) ([]dto.Coordinates, error) {
	//Get all step distances in lowest unit (metres)
	distances := getStepDistances(requestID, req)
	
	//Get GCD
	minInterval := getMinimumInterval(requestID, distances)

	//Get Coordinates
	coordinates := getCoordinates(requestID, req, minInterval)

	return coordinates, nil
}

func getStepDistances(requestID string, req model.GoogleClientGetDirectionsResponse) []int {
	var distances []int
	
	for _, route := range *req.Routes {
		for _, leg := range *route.Legs {
			for _, step := range *leg.Steps {
				dist := *step.Distance.Value
				distances = append(distances, dist)
			}
		}
	}

	return distances
}

func getMinimumInterval(requestID string, distances []int) int {
	gcd := distances[0]

	if len(distances) == 1 {
		return gcd
	}

	for _, val := range distances {
		gcd = common.ComputeGCD(gcd, val)
	}

	return gcd
}

func getCoordinates (requestID string, req model.GoogleClientGetDirectionsResponse, minInterval int) []dto.Coordinates {
	var coordinates []dto.Coordinates

	for _, route := range *req.Routes {
		for _, leg := range *route.Legs {
			for _, step := range *leg.Steps {
				distance := step.Distance.Value
				start := step.StartLocation
				end := step.EndLocation
				coord := getCoordinatesForStep(requestID, distance, start, end, minInterval)
				coordinates = append(coordinates, coord...)
			}
		}
	}	

	return coordinates
}

func getCoordinatesForStep(requestID string, distance *int, start *model.LocationData, end *model.LocationData, minInterval int) []dto.Coordinates {
	var coordinates []dto.Coordinates
	intervals := (*distance) / minInterval
	stepInterval := float64(1 / intervals)

	phi1, phi2, th1, th2, dph, dth := getValuesInRadians(start, end)

	for i := 1; i < intervals; i++ {
		stepInterval = stepInterval * float64(i)

		a := (math.Sin(dph / 2) * math.Sin(dph / 2)) + (math.Cos(phi1) * math.Cos(phi2) * math.Sin(dth / 2) * math.Sin(dth / 2))
		delta := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))

		A := math.Sin((1 - stepInterval) * delta) / math.Sin(delta)
		B := math.Sin(stepInterval * delta) / math.Sin(delta)

		x := (A * math.Cos(phi1) * math.Cos(th1)) + (B * math.Cos(phi2) * math.Cos(th2))
		y := (A * math.Cos(phi1) * math.Sin(th1)) + (B * math.Cos(phi2) * math.Sin(th2))
		z := (A * math.Sin(phi1)) + (B * math.Sin(phi2))

		phi3 := math.Atan2(z, math.Sqrt((x * x) + (y * y)))
		th3 := math.Atan2(y, x)

		lat := common.ToDegrees(phi3)
		lon := common.ToDegrees(th3)

		coord := dto.Coordinates{
			Latitude: lat,
			Longitude: lon,
		}

		coordinates = append(coordinates, coord)

	}

	return coordinates


}

func getValuesInRadians(start *model.LocationData, end *model.LocationData) (float64, float64, float64, float64, float64, float64) {
	ph1 := common.ToRadians(*start.Latitude)
	ph2 := common.ToRadians(*end.Latitude)
	th1 := common.ToRadians(*start.Longitude)
	th2 := common.ToRadians(*end.Longitude)

	dph := ph2 - ph1
	dth := th2 - th1

	return ph1, ph2, th1, th2, dph, dth
}
