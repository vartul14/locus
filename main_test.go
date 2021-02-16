package main

import (
	"testing"
)
func TestAddBill(t *testing.T) {
	impl := SplitwiseImpl{}

	paidBy := map[int]float64{
		1: 200,
	}

	owedBy := map[int]float64 {
		2: 100,
		3: 100,
	}

	spType := Exact

	impl.AddBill(paidBy , owedBy , 0, spType)

}