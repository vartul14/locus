package google_client

import (
	"fmt"
	"vartul14/locus/minion/google_client/enums"
	"vartul14/locus/minion/google_client/model"

	"github.com/go-resty/resty"
	"encoding/json"

)

type GoogleClient struct {
	HostAndPort string
	Headers     map[string]string
}

type Client interface {
	GetDirections(requestID string, req map[string]string) (model.GoogleClientGetDirectionsResponse, int, error)
}

func (client *GoogleClient) GetDirections(requestID string, req map[string]string) (model.GoogleClientGetDirectionsResponse, int, error) {
	fmt.Printf("API Call Client = %v | API = %v | RequestID = %v | RequestData = %v", enums.GoogleClientTag, enums.GetDirectionsAPITag, requestID, req)
	url := (client.HostAndPort) + enums.GetDirectionsURI

	restyClient := resty.New()

	resp, err := restyClient.R().
		SetQueryParams(req).
		SetHeaders(client.Headers).
		Get(url)
	
	getDirectionsResponse := model.GoogleClientGetDirectionsResponse{}
	json.Unmarshal(resp.Body(), &getDirectionsResponse)

	return getDirectionsResponse, resp.StatusCode(), err
}
