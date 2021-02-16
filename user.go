package main

type Splitwise interface {
	AddBill(paidBy map[int]float64, owedBy map[int]float64, groupID int, splitType SplitType)
	GetGroupBalances(userID int, groupID int) map[int]Balance
	GetTotalBalances(userID int) map[int]Balance
}

type SplitwiseImpl struct {

}

func (impl *SplitwiseImpl) AddBill(paidBy map[int]float64, owedBy map[int]float64, groupID int, splitType SplitType) {
	split := GetSplitImplByType(splitType).GetSplit(paidBy, owedBy)

	GetDaoImpl(InMemory).UpdateUser(split)
	GetDaoImpl(InMemory).UpdateGroup(split, groupID)
}

func (impl *SplitwiseImpl) GetGroupBalances(userID int, groupID int) map[int]Balance {
	balances := GetDaoImpl(InMemory).GetGroupBalances(userID, groupID)
	return balances
}

func (impl *SplitwiseImpl) GetToalBalances(userID, groupID int) map[int]Balance {
	balances := GetDaoImpl(InMemory).GetTotalBalances(userID)
	return balances
}


