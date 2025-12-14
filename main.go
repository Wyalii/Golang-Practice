package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	intSlice := intArr[1:6]
	fmt.Printf("the length is %v, with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println()
	intSlice = append(intSlice, 11, 12, 13)
	fmt.Printf("the length is %v, with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println()
	fmt.Println(intArr)
	fmt.Println(intSlice)
	fmt.Println(intArr)

}

// func printMe(printVaule string) {
// 	fmt.Println(printVaule)
// 	var result, remainder, err = intDivision(8, 10)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("the result of divison is %v with remainder %v\n", result, remainder)
// 	switch remainder {
// 	case 0:
// 		fmt.Println("the division was succesfull, no remainder!")
// 	case 1, 2:
// 		fmt.Println("clean division was really close!")
// 	default:
// 		fmt.Println("the division was not close!")
// 	}

// }
// func intDivision(numerator int, denominator int) (int, int, error) {
// 	var err error
// 	if denominator == 0 {
// 		err = errors.New("cant divide by zero")
// 		return 0, 0, err
// 	}
// 	var result int = numerator / denominator
// 	var remainder int = numerator % denominator
// 	return result, remainder, err
// }

// // func floatDivision(numerator float64, denominator float64) float64 {
// // 	var result float64 = numerator / denominator
// // 	return result
// // }
