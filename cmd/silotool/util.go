package main

import (
	"reflect"
	"encoding/hex"
	"fmt"
)

// this function should only be used on structs containing only byte slices
func printStructOfBytesAsHex(stct interface{}) {
	v := reflect.ValueOf(stct)
	t := v.Type()

	for i:=0; i<v.NumField(); i++ {
		name := t.Field(i).Name
		value := hex.EncodeToString(v.Field(i).Interface().([]byte))
		fmt.Println(name, value)
	}
}