package custompackage

import (
	"fmt"
)

func init(){
	fmt.Println("Initializing rectangle package.")
}

//functions with first letter capital are public functions
func Area(width, length float64) float64 {
	area := width * length
	return area
}