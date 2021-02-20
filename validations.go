package main

import (
	"reflect"
)

type Validations struct {
	ColType reflect.Type
	MinVal interface{}
	MaxVal interface{}
	IsMandatory bool
}
