package main

import (
	"fmt"
	"math"
)

func main(){
	
	//New variable without initial value
	var age int
	fmt.Println("Initial age is equal to", age)
	age = 23
	fmt.Println("Age is now equal to", age)

	//New variable with initial value
	var year = 2018
	fmt.Println("Initial year is equal to", year)

	//Two new variables of the same type with initial values
	var firstName, lastName = "Matt", "Cole"
	fmt.Println("my name is", firstName, lastName)

	//Two new variables of different types
	var(
		personName = "Jim"
		personAge = 26
	)
	fmt.Println(personName,"is",personAge,"years old.")

	//Declare variables without 'var' keyword
	street, number := "W Princess Anne", 625
	fmt.Println("The address is", number, street)

	a, b := 3.0, 10.0
	c := math.Min(a, b)
	fmt.Println("The min of", a,"and", b, "is", c)
}