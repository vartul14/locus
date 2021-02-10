package api

import (
	"encoding/json"

	"fmt"

	"net/http"



	"github.com/golang/gddo/httputil/header"
)

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			fmt.Printf("Content-Type header is not application/json")
			return fmt.Errorf("Error parsing request")
		}
	}

	pathParams := make(map[string]interface{})
	pathParams["request-id"] = r.URL.Query().Get("request-id")
	pathParams["start-lat"] = r.URL.Query().Get("start-lat")
	pathParams["start-lon"] = r.URL.Query().Get("start-lon")
	pathParams["end-lat"] = r.URL.Query().Get("end-lat")
	pathParams["end-lon"] = r.URL.Query().Get("end-lon")

	data, _ := json.Marshal(pathParams)
	fmt.Printf(string(data))
	_ = json.Unmarshal(data, dst)
	return nil
}
