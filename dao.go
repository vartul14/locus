package main

type StorageInterface interface {
	UpdateUser(split map[int]map[int]Balance)
	UpdateGroup(split map[int]map[int]Balance, groupID int)
	GetGroupBalances(userID int, groupID int) map[int]Balance
	GetTotalBalances(userID int) map[int]Balance
}

type InMemoryImpl struct {}

func GetDaoImpl(storageType StorageType) StorageInterface {
	switch storageType {
	case InMemory:
		return &InMemoryImpl{}
	}

	return nil
}

func (impl *InMemoryImpl) UpdateUser (split map[int]map[int]Balance) {
	for k, u := range split {
		for dep, b := range u {
			amount := b.Amount
			userData := Users[k]
			currAmount := userData.Balances[dep].Amount
			tAmount := currAmount + amount
			am := userData.Balances[dep]
			am.Amount = tAmount
			Users[k] = userData
		}
	}
}

func (impl *InMemoryImpl) UpdateGroup(split map[int]map[int]Balance, groupID int) {
	for k, u := range split {
		for dep, b := range u {
			group := Groups[groupID]
			amount := b.Amount
			currAmount := group.Balances[k][dep]
			tAmount := currAmount.Amount + amount
			am := group.Balances[k][dep]
			am.Amount = tAmount
			Groups[groupID] = group
		}
	}
}

func (impl *InMemoryImpl) GetGroupBalances(userID int, groupID int) map[int]Balance{
	
	group := Groups[groupID]
	return group.Balances[userID]
}

func (impl *InMemoryImpl) GetTotalBalances(userID int) map[int]Balance{
	
	user := Users[userID]
	return user.Balances
}