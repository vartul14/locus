package dto

type GetTransactionsRequest struct {
	RequestID     string      `json:"request-id"`
	TransactionID int `json:"transaction-id"`
}

type GetTransactionsResponse struct {
	RequestID  string        `json:"request-id"`
	Amount float64 `json:"amount"`
	Type   string `json:"type"`
	ParentID    int `json:"parent-id"`
}

