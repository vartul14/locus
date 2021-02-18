package api

import (
	"encoding/json"

	"fmt"

	"net/http"
	"strconv"



	"github.com/golang/gddo/httputil/header"
)

type Error struct {
	ErrorMessage string
}

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
	pathParams["transaction-id"], _ = strconv.Atoi(r.URL.Query().Get("transaction-id"))
	pathParams["transaction-type"] = r.URL.Query().Get("transaction-type")

	data, _ := json.Marshal(pathParams)
	//fmt.Println(string(data))
	_ = json.Unmarshal(data, dst)
	return nil
}

func sendErrorResponse(w http.ResponseWriter, logicErr error) {
	err := Error{
		ErrorMessage: logicErr.Error(),
	}

	jsonRes, _ := json.Marshal(err)
	fmt.Fprintf(w, string(jsonRes))
}
