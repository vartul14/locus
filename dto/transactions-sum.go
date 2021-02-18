package dto

type GetTransactionsSumRequest struct {
	RequestID     string      `json:"request-id"`
	TransactionID int `json:"transaction-id"`
}

type GetTransactionsSumResponse struct {
	RequestID  string        `json:"request-id"`
	Amount float64 `json:"amount"`
}

