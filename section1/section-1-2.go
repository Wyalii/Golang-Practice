package main

import (
	"fmt"
	"os"
)

// this is a package level variable that is accesable in every file as long as it has package main and its on same folder
var thisIsMyPackageLevelVariable string = "hello world"

// this is another type of loop that uses range keyword, it produces key value pairs on each iteration,
// one for index and one for values of the element on that index.
func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		fmt.Println("index:", i)
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	//this declaration is mostly used in functions not on package level variables
	string1 := ""

	//this one relies on default initalization wich for strings is empty ""
	var string2 string

	//this one is rarely used except when declaring multiple variables
	var string3 = ""

	//idk about this one but it looks clean to me :)
	var string4 string = ""

	fmt.Println(string1, string2, string3, string4)

	//here im using function that is declared in other file that uses package lvl variable declared in this file cool :)
	printPackageLevelVariable()
}
