package logic

import (
	"vartul14/locus/dao"
	"vartul14/locus/dto"

)

func GetTransactions(request dto.GetTransactionsRequest) (*dto.GetTransactionsResponse, error) {
	transactionData := dao.InMemoryDB.GetTransaction(request.TransactionID)

	resp := &dto.GetTransactionsResponse{
		RequestID: request.RequestID,
		Amount: transactionData.Amount,
		Type: transactionData.TransactionType,
		ParentID: transactionData.ParentID,
	}

	return resp, nil
}


