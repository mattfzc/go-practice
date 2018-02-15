package main

import "fmt"

func main(){
	number := 4

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("number out of range")
	}

	letter := "a"

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("This is a vowel")
	default:
		fmt.Println("not a vowel")
	}

	number = 5
	switch{
	case number < 2:
		fmt.Println("less than 2")
	case number < 6:
		fmt.Println("less than 6")
	case number < 8:
		fmt.Println("less than 8")
	}

	switch {
	case number < 2:
		fmt.Println("less than 2")
		fallthrough
	case number < 6:
		fmt.Println("less than 6")
		fallthrough
	case number < 8:
		fmt.Println("less than 8")
	}
}