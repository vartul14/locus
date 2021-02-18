package dto

type GetTransactionsFromTypeRequest struct {
	RequestID     string      `json:"request-id"`
	Type string `json:"transaction-type"`
}

type GetTransactionsFromTypeResponse struct {
	RequestID  string        `json:"request-id"`
	TransactionIDs []int `json:"transaction-ids"`
}