package main

import (
	"testing"
	"reflect"
)

func TestCreateTable(t *testing.T) {
	impl := InMemoryImpl{}
	impl.Init()


	validationMap:= make(map[string]Validations)
	val := Validations {
		ColType: reflect.TypeOf(1),
		MinVal: 0,
		MaxVal: 100,
		IsMandatory: true,
	}
	validationMap["id"] = val

	err := impl.AddTable("table1", validationMap)
	if err != nil {
		t.Errorf("Error adding table")
	}

}

func TestInsertRow(t *testing.T) {
	impl := InMemoryImpl{}
	impl.Init()


	validationMap:= make(map[string]Validations)
	val := Validations {
		ColType: reflect.TypeOf(1),
		MinVal: 0,
		MaxVal: 100,
		IsMandatory: true,
	}

	row := make(map[string]interface{})
	row["id"] = 1
	validationMap["id"] = val

	err := impl.AddTable("table1", validationMap)
	errInsert := impl.InsertRow("table1", row)
	if err != nil {
		t.Errorf("Error adding table")
	}
	if errInsert != nil {
		t.Errorf("Error inserting row")
	}

}

func TestPrintRows(t *testing.T) {
	impl := InMemoryImpl{}
	impl.Init()


	validationMap:= make(map[string]Validations)
	val := Validations {
		ColType: reflect.TypeOf(1),
		MinVal: 0,
		MaxVal: 100,
		IsMandatory: true,
	}

	row := make(map[string]interface{})
	row["id"] = 1
	validationMap["id"] = val

	err := impl.AddTable("table1", validationMap)
	errInsert := impl.InsertRow("table1", row)
	_, errPrint := impl.PrintAllRows("table1")
	if err != nil {
		t.Errorf("Error adding table")
	}
	if errInsert != nil {
		t.Errorf("Error inserting row")
	}
	if errPrint != nil {
		t.Errorf ("Error printing the rows")
	}
	
}

func TestFilterRows(t *testing.T) {
	impl := InMemoryImpl{}
	impl.Init()


	validationMap:= make(map[string]Validations)
	val := Validations {
		ColType: reflect.TypeOf(1),
		MinVal: 0,
		MaxVal: 100,
		IsMandatory: true,
	}
	validationMap["id"] = val

	row1 := make(map[string]interface{})
	row1["id"] = 1
	

	row2 := make(map[string]interface{})
	row2["id"] = 2
	

	filter := Filter {
		columnName: "id",
		value: 2,
	}
	filterParams := []Filter{}
	filterParams = append(filterParams, filter)

	_ = impl.AddTable("table1", validationMap)
	errInsert := impl.InsertRow("table1", row1)
	_ = impl.InsertRow("table1", row2)
	_, errPrint := impl.PrintAllRows("table1")
	_, errFilter := impl.FilterRows("table1", filterParams)
	// if err != nil {
	// 	t.Errorf("Error adding table")
	// }
	if errInsert != nil {
		t.Errorf("Error inserting row")
	}
	if errPrint != nil {
		t.Errorf ("Error printing the rows")
	}
	if errFilter != nil {
		t.Errorf("Error filtering")
	}
	// if matchingRows[0]["id"] != 2 {
	// 	t.Errorf("incorrect filtering")
	// }
}

func TestValidations(t *testing.T) {
	impl := InMemoryImpl{}
	impl.Init()


	validationMap:= make(map[string]Validations)
	valID := Validations {
		ColType: reflect.TypeOf(1),
		MinVal: 0,
		MaxVal: 100,
		IsMandatory: true,
	}

	valName := Validations {
		ColType: reflect.TypeOf("vartul"),
		//MinVal: 0,
		MaxVal: 10,
		IsMandatory: false,
	}

	row := make(map[string]interface{})
	validationMap["id"] = valID
	validationMap["name"] = valName

	row["id"] = 2
	//row["name"] = 

	filter := Filter {
		columnName: "id",
		value: 1,
	}
	filterParams := []Filter{}
	filterParams = append(filterParams, filter)

	err := impl.AddTable("table1", validationMap)
	errInsert := impl.InsertRow("table1", row)
	_, errPrint := impl.PrintAllRows("table1")
	_, errFilter := impl.FilterRows("table1", filterParams)
	if err != nil {
		t.Errorf("Error adding table")
	}
	if errInsert != nil {
		t.Errorf("should get error while inserting")
	}
	if errPrint != nil {
		t.Errorf ("Error printing the rows")
	}
	if errFilter != nil {
		t.Errorf("Error filtering")
	}
	// if matchingRows[0]["id"] != 1 {
	// 	t.Errorf("incorrect filtering")
	// }
}