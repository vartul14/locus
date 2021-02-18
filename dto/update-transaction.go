package dto

type UpdateTransactionsRequest struct {
	RequestID     string      `json:"request-id"`
	TransactionID int `json:"transaction-id"`
	Amount float64 `json:"amount"`
	Type   string `json:"type"`
	ParentID    int `json:"parent-id"`
}

type UpdateTransactionsResponse struct {
	RequestID  string        `json:"request-id"`
	Status     string        `json:"status"`
}