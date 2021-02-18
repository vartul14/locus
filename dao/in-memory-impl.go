package dao


type InMemoryImpl struct {

	//Multiple DB represent DB indexing on each columns
	InMemoryDBByID map[int]TransactionData
	InMemoryDBByType map[string][]int
	InMemoryDBByUltimateParent map[int][]int
}

type TransactionData struct {
	ID int
	Amount float64
	TransactionType string
	ParentID int
	UltimateParentID int
}

var InMemoryDB InMemoryImpl

func init () {
	InMemoryDB.InMemoryDBByID = make(map[int]TransactionData)
	InMemoryDB.InMemoryDBByType = make(map[string][]int)
	InMemoryDB.InMemoryDBByUltimateParent = make(map[int][]int)
}

func (impl *InMemoryImpl) UpdateTransaction(transactionID int, transactionData TransactionData) {
	db := (*impl).InMemoryDBByID
	dbByType := (*impl).InMemoryDBByType
	dbByUltimateParent := (*impl).InMemoryDBByUltimateParent

	transactionData.ID = transactionID
	db[transactionID] = transactionData

	li := dbByType[transactionData.TransactionType]
	li = append(li, transactionID)
	dbByType[transactionData.TransactionType] = li

	list := dbByUltimateParent[transactionData.UltimateParentID]
	list = append(list, transactionID)
	dbByUltimateParent[transactionData.UltimateParentID] = list
}

func (impl *InMemoryImpl) GetTransaction (transactionID int) TransactionData {
	db :=(*impl).InMemoryDBByID
	
	transactionData := db[transactionID]
	return transactionData
}
	
func (impl *InMemoryImpl) GetTransactionsByType (transactionType string) []int {
	db := (*impl).InMemoryDBByType

	transactionList := db[transactionType]
	return transactionList
}
	
func (impl *InMemoryImpl) GetTransactionsByUltimateParentID (ultimateParentID int) []int {

	dbUltimateParentID := (*impl).InMemoryDBByUltimateParent

	listTransactionID := dbUltimateParentID[ultimateParentID]
	return listTransactionID

}
