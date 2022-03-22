package main

import (
	"fmt"
	"reflect"
)

// func main() {
// 	var number = 23
// 	var reflectValue = reflect.ValueOf(number)

// 	fmt.Println("Type of variable:", reflectValue.Type())

// 	if reflectValue.Kind() == reflect.Int {
// 		fmt.Println("variable value :", reflectValue.Int())
// 	}
// }

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Jalan", 20}

	var sampleType = reflect.TypeOf(person)
	var valueType = reflect.ValueOf(person)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Type)
	fmt.Println(valueType.Field(0).Interface())
}
