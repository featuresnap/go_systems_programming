package main

import (
	"fmt"
	"reflect"
)

func main() {
	type t1 int
	type t2 int

	i1 := t1(1)
	i2 := t2(1)

	fmt.Println("i1:", i1)
	fmt.Println("i2:", i2)

	//types
	x1 := reflect.TypeOf(i1)
	x2 := reflect.TypeOf(i2)

	fmt.Println("reflect.TypeOf(i1):", x1)
	fmt.Println("reflect:TypeOf(i2):", x2)

	//values
	fmt.Println("reflect.ValueOf(i1):", reflect.ValueOf(i1))
	fmt.Println("reflect:ValueOf(i2):", reflect.ValueOf(i2))

	//to compare two values, you need to call their Interface() methods.
	fmt.Println("Reflected values equal?", reflect.ValueOf(i1) == reflect.ValueOf(i2))
	fmt.Println("Reflected interface equal?", reflect.ValueOf(i1).Interface() == reflect.ValueOf(i2).Interface())

}
