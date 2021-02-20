package main

import(
	"reflect"
)

func checkMandatory(isMandatory bool, colName string, row map[string]interface{}) bool {
	if isMandatory {
		if _, ok := row[colName]; !ok {
			return false
		}
	}

	return true
}

func checkType(colType reflect.Type, colName string, row map[string]interface{}) bool {
	val, ok := row[colName]
	if ok && (reflect.TypeOf(val) != colType) {
		return false
	} 
	return true
}

func checkRange(min interface{}, max interface{}, colName string, row map[string]interface{}) bool {
	val, ok := row[colName]

	if !ok {
		return true
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		if min != nil && val.(int) < min.(int) {
			return false
		}
		if max != nil && val.(int) > max.(int) {
			return false
		}
	case reflect.String:
		s := val.(string)
		if len(s) > max.(int) {
			return false
		}
	}
	
	

	return true
}