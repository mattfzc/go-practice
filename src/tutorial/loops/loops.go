package main

import "fmt"

func main(){

	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}

		if i % 2 == 0 {
			continue
		}
		fmt.Printf(" %d", i)
	} 

	index := 1

	for index <= 10{
		fmt.Printf(" %d", index)
		index++
	}

	for {
		if(index == 20){
			break
		}
		fmt.Printf(" %d", index)
		index++

	}
}