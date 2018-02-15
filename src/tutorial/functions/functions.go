package main

import "fmt"

func main(){

	price, amount := 6, 3
	totalPrice := calculateBill(price, amount)
	fmt.Printf("The total price %d meals at $%d a meal is $%d", amount, price, totalPrice)
	
	width, length := 3.5, 7.0
	area, perimiter := getAreaAndPerimeter(width, length)
	fmt.Printf("\nWhen width is %f and length is %f, area is %f and perimiter is %f", width,length,area,perimiter)
}
//Functions with first letter lower case are private functions
func calculateBill(price int, amount int) int {
	totalPrice := price * amount
	return totalPrice
}

func getAreaAndPerimeter(width, length float64) (float64, float64) {
	area := width * length
	perimiter := (width + length) * 2
    return area, perimiter
}