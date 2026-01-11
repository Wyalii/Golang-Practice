package main

import (
	"fmt"
)

// using package level variable that is declared in section-1-2.go here
func printPackageLevelVariable() {
	fmt.Println(thisIsMyPackageLevelVariable)
}
