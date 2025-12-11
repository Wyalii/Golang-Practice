package main

import (
	"fmt"
)

func main() {
	printMe("ASKJKAJSD")
}
func printMe(printVaule string) {
	fmt.Println(printVaule)
	var result, remainder int = intDivision(8, 0)

	fmt.Println(result, remainder)

}
func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
func floatDivision(numerator float64, denominator float64) float64 {
	var result float64 = numerator / denominator
	return result
}
