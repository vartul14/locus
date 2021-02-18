package api

import (
	"fmt"
	"net/http"
	"vartul14/locus/dto"
	"vartul14/locus/logic"
	"encoding/json"
)


func UpdateTransactions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Update Transactions API")

	var request dto.UpdateTransactionsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	updateTransactionsResponse, logicErr := logic.UpdateTransactions(request)
	if logicErr != nil {
		fmt.Printf("Error in logic to update transactions | Error = %v", logicErr)
		sendErrorResponse(w, logicErr)
	} else {
		jsonResp, _ := json.Marshal(updateTransactionsResponse)
		fmt.Fprintf(w, string(jsonResp))
	}
	
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Get Transactions API")

	var request dto.GetTransactionsRequest
	err := decodeJSONBody(w, r, &request)
	if err != nil {
		fmt.Printf("Error in decoding the request | Error = %v", err.Error())
		return
	}

	getTransactionsResponse, logicErr := logic.GetTransactions(request)
	if logicErr != nil {
		fmt.Printf("Error in logic to get transaction | Error = %v", logicErr.Error())
		sendErrorResponse(w, logicErr)
	} else {
		jsonResp, _ := json.Marshal(getTransactionsResponse)
		fmt.Fprintf(w, string(jsonResp))
	}
	
}

func GetTransactionsFromTypes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Get Transactions From Types API")

	var request dto.GetTransactionsFromTypeRequest
	err := decodeJSONBody(w, r, &request)
	if err != nil {
		fmt.Printf("Error in decoding the request | Error = %v", err.Error())
		return
	}

	getTransactionsResponse, logicErr := logic.GetTransactionsFromType(request)
	if logicErr != nil {
		fmt.Printf("Error in logic to get transaction | Error = %v", logicErr.Error())
		sendErrorResponse(w, logicErr)
	} else {
		jsonResp, _ := json.Marshal(getTransactionsResponse)
		fmt.Fprintf(w, string(jsonResp))
	}
}

func GetTransactionsSum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Get Transactions From Types API")

	var request dto.GetTransactionsSumRequest
	err := decodeJSONBody(w, r, &request)
	if err != nil {
		fmt.Printf("Error in decoding the request | Error = %v", err.Error())
		return
	}

	getTransactionsResponse, logicErr := logic.GetTransactionsSum(request)
	if logicErr != nil {
		fmt.Printf("Error in logic to get transaction | Error = %v", logicErr.Error())
		sendErrorResponse(w, logicErr)
	} else {
		jsonResp, _ := json.Marshal(getTransactionsResponse)
		fmt.Fprintf(w, string(jsonResp))
	}
}
