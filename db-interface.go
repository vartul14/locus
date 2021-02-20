package main 

type DB interface {
	AddTable(tableName string, validations map[string]Validations) error
	DeleteTable(tableName string) error
	InsertRow(tableName string, row map[string]interface{}) error
	PrintAllRows(tableName string) (map[int]map[string]interface{}, error)
	FilterRows(tableName string, filterParams []Filter) ([]map[string]interface{}, error)
}

