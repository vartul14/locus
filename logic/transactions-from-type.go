package logic

import (
	"vartul14/locus/dao"
	"vartul14/locus/dto"

)

func GetTransactionsFromType(request dto.GetTransactionsFromTypeRequest) (*dto.GetTransactionsFromTypeResponse, error) {
	transactionList := dao.InMemoryDB.GetTransactionsByType(request.Type)

	resp := &dto.GetTransactionsFromTypeResponse{
		RequestID: request.RequestID,
		TransactionIDs: transactionList,
	}

	return resp, nil
}
