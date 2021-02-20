package main

import (
	"fmt"
)

type InMemoryImpl struct {
	DB map[string]map[int]map[string]interface{}
	Validation map[string]map[string]Validations     //check how to add validations//can move to input layer
	RowNumber map[string]int
}

func  (i *InMemoryImpl) Init() {
	(*i).DB = make(map[string]map[int]map[string]interface{})
	(*i).Validation = make(map[string]map[string]Validations)
	(*i).RowNumber = make(map[string]int)
}

func (i *InMemoryImpl) AddTable(tableName string, validations map[string]Validations) error {
	db := (*i).DB
	val := (*i).Validation
	rNum := (*i).RowNumber

	//Check if table already exists
	if ok := i.checkIfTableExists(tableName); ok {
		return fmt.Errorf("Table already exists")
	}


	//Create a new table in the map
	db[tableName] = make(map[int]map[string]interface{})
	rNum[tableName] = 0


	//Add the validation for the table in the validations map
	val[tableName] = validations

	return nil
}

func (i *InMemoryImpl) DeleteTable(tableName string) error {
	db := (*i).DB
	val := (*i).Validation
	rNum := (*i).RowNumber

	//check if the table exists in the db
	if ok := i.checkIfTableExists(tableName); !ok {
		return fmt.Errorf("Table does not exists, cannot delete")
	}

	//Delete the table from the DB
	delete(db, tableName)

	//Delete the validations form the validation struct 
	delete(val, tableName)

	delete(rNum, tableName)

	return nil
}

func (i *InMemoryImpl) InsertRow(tableName string, row map[string]interface{}) error {
	db := (*i).DB
	val := (*i).Validation
	rNum := (*i).RowNumber

	//Check if table exists
	if ok := i.checkIfTableExists(tableName); !ok {
		return fmt.Errorf("Table does not exist")
	}

	//Check for input validations
	if inputCheck := i.checkForInputValidations(val[tableName], row); !inputCheck {
		return fmt.Errorf("Input validation check failed, cannot insert row")
	}

	//Update the table with a new row in DB
	nextRowID := rNum[tableName]
	rNum[tableName] = nextRowID + 1
	tableData := db[tableName]
	tableData[nextRowID] = row

	return nil
}

func (i *InMemoryImpl) PrintAllRows(tableName string) (map[int]map[string]interface{}, error) {
	db := (*i).DB
	
	//Check if table exists
	if ok := i.checkIfTableExists(tableName); !ok {
		return nil, fmt.Errorf("Table does not exist")
	}

	//Return all rows from the table
	tableData := db[tableName]
	fmt.Println(tableData)
	return tableData, nil


}

func (i *InMemoryImpl) FilterRows(tableName string, filterParams []Filter) ([]map[string]interface{}, error) {
	db := (*i).DB
	
	//Check if table exists
	if ok := i.checkIfTableExists(tableName); !ok {
		return nil, fmt.Errorf("Table does not exist")
	}

	tableData := db[tableName]

	//Get matching rows in the table
	matchingRows := i.getMatchingRows(tableData, filterParams)
	fmt.Println(matchingRows)

	return matchingRows, nil
}

func (i *InMemoryImpl) checkIfTableExists (tableName string) bool {
	db := (*i).DB
	if _, ok := db[tableName]; !ok {
		return false
	}

	return true
}

func (i *InMemoryImpl) getMatchingRows(tableData map[int]map[string]interface{}, filterParams []Filter) []map[string]interface{} {
	matchingRows := make([]map[string]interface{}, 0)
	
	for _, row := range tableData {
		for _, value := range filterParams {
			col := value.columnName
			filterVal := value.value

			if val, ok  := row[col]; ok {
				if filterVal == val {
					matchingRows = append(matchingRows, row)
				} else {
					break
				}
			} else {
				break
			}
		}
	}

	return matchingRows
}


func (i *InMemoryImpl) checkForInputValidations(validations map[string]Validations, row map[string]interface{}) bool {
	for colName, validate := range validations {
		if mandatoryCheck := checkMandatory(validate.IsMandatory, colName, row); !mandatoryCheck {
			return false
		}
		if typeCheck := checkType(validate.ColType, colName, row); !typeCheck {
			return false
		}
		if rangeCheck := checkRange(validate.MinVal, validate.MaxVal, colName, row); !rangeCheck {
			return false
		}
		
	}

	return true
}
