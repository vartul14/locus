package main

type SplitInterface interface {
	GetSplit(paidBy map[int]float64, owedBy map[int]float64) map[int]map[int]Balance
}

func GetSplitImplByType(splitType SplitType) SplitInterface {
	switch splitType {
	case Exact:
		return &ExactImpl{}
	
	}

	return nil
}

type ExactImpl struct {}

func (impl *ExactImpl) GetSplit(paidBy map[int]float64, owedBy map[int]float64) map[int]map[int]Balance {
	totalAmountPaid := 0
	totalUsers := 0

	for val := range paidBy {
		totalAmountPaid = totalAmountPaid + val
	}

	totalUser := len(owedBy)
	owes := float64(totalAmountPaid / totalUsers)

	m := make(map[int]map[int]Balance)
	for p, _ := range paidBy {
		for x, val := range owedBy {
			m[x][p] = Balance{
				Amount: float64(-1) * (owes),
			}

			m[p][x] = Balance {
				Amount: float64(-1) * (owes),
			}
		}
		
	}

	return m

}


type PercantageImpl struct {}

func (impl *PercantageImpl) GetSplit(paidBy map[int]int, owedBy map[int]int) map[int]map[int]Balance {
	m := make(map[int]map[int]Balance)
	return m

}