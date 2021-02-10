package api

import (
	"fmt"
	"net/http"
	"vartul14/locus/dto"
	"vartul14/locus/logic"
)

func GetDirections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetDirections API")

	var request dto.GetDirectionsRequest
	err := decodeJSONBody(w, r, &request)
	if err != nil {
		fmt.Printf("Error in decoding the request | Error = %v", err.Error())
		return
	}

	getDirectionsResponse, logicErr := logic.GetDirections(request)
	if logicErr != nil {
		fmt.Printf("Error in logic to get directions | Error = %v", err.Error())
	}
	fmt.Println(getDirectionsResponse)
}
