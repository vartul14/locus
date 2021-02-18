package dao


type DaoInterface interface {
	UpdateTransaction(transactionID int, transactionData TransactionData)
	GetTransaction (transactionID int) TransactionData
	GetTransactionsByType (transactionType string) TransactionData
	GetTransactionsByUltimateParentID (ultimateParentID int) []TransactionData
}