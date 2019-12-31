package main

import "reflect"

// CreatePointerSlice create slice of struct/ptrStruct with len and cap and return result as ptrSlice.
// To create empty slice set len and cap = `0`.
// e.g: `structOrPtrStruct = Object{}, it will return *[]Object`
func CreatePointerSlice(structOrPtrStruct interface{}, len, cap int) (ptrSlice interface{}) {
	typ := reflect.TypeOf(structOrPtrStruct)
	tmpSlice := reflect.MakeSlice(reflect.SliceOf(typ), len, cap)
	slice := reflect.New(tmpSlice.Type())
	ind := reflect.Indirect(slice)
	slice.Elem().Set(ind)

	return slice.Interface()
}
