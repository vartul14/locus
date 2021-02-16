package main

type SplitType string
const (
	Exact SplitType = "Exact"
	Percentage SplitType = "Percentage"
)

type StorageType string
const (
	InMemory StorageType = "InMemory"
	DB StorageType = "DB"
)