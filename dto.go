package main

type User struct {
	Name string
	UserID int
	Balances map[int]Balance
}

type Group struct {
	GroupName string
	GroupID int
	Balances map[int]map[int]Balance
}
type Balance struct {
	Amount float64
}

type UserData map[int]User
type GroupData map[int]Group

var Users UserData
var Groups GroupData
