package main

import (
	"fmt"
	"unsafe"
)
func main(){

	a, b := true, false
	fmt.Println("A is", a, "and B is",b)
	c := a || b
	fmt.Println("C is",c)
	d := a && b
	fmt.Println("D is",d)
	var numberA int = 89
	numberB := 95
	fmt.Printf("The type of A is %T and the size is %d. The type of B is %T and the size is %d.", numberA, unsafe.Sizeof(numberA), numberB, unsafe.Sizeof(numberB))
	decimalValue := 16.8
	integer := int(decimalValue)
	fmt.Println("The decimal",decimalValue,"converted to an integer is",integer)
}