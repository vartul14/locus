package logic

import (
	"vartul14/locus/dao"
	"vartul14/locus/dto"

)


func GetTransactionsSum (request dto.GetTransactionsSumRequest) (*dto.GetTransactionsSumResponse, error) {
	transactionData := dao.InMemoryDB.GetTransaction(request.TransactionID)
	listTransactionID := dao.InMemoryDB.GetTransactionsByUltimateParentID(transactionData.UltimateParentID)
	
	sum := 0.0
	for _, val := range listTransactionID {
		transaction := dao.InMemoryDB.GetTransaction(val)
		sum = sum + transaction.Amount
	}

	resp := &dto.GetTransactionsSumResponse{
		RequestID: request.RequestID,
		Amount: sum,
	}

	return resp, nil
}
