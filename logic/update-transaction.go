package logic

import (
	"vartul14/locus/dao"
	"vartul14/locus/dto"

	"fmt"

)


func UpdateTransactions(request dto.UpdateTransactionsRequest) (*dto.UpdateTransactionsResponse, error) {
	
	parentData := dao.TransactionData{
		ID: request.TransactionID,
	}
	if request.ParentID != 0 {
		parentData = dao.InMemoryDB.GetTransaction(request.ParentID)
	}

	if parentData.ID == 0 {
		return &dto.UpdateTransactionsResponse{}, fmt.Errorf("Parent transaction does not exist")
	}
	
	updateRequest := dao.TransactionData {
		Amount: request.Amount,
		TransactionType: request.Type,
		ParentID: request.ParentID, 
		UltimateParentID: parentData.ID,
	}

	dao.InMemoryDB.UpdateTransaction(request.TransactionID, updateRequest)

	resp := &dto.UpdateTransactionsResponse{
		RequestID: request.RequestID,
		Status: "success",
	}

	return resp, nil
}
