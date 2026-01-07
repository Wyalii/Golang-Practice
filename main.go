package main

import (
	"fmt"
	"os"
)

func main() {
	// concatenation of string os package and command line arguments in args.
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// looks like golang only has 1 loop which is for loop that that different forms//

	//this one is while loop lol//
	var randomNumberHere int16
	for randomNumberHere < 16 {
		println(randomNumberHere)
		randomNumberHere++
	}
	//loop with no initialization, condition or post is just infinite loop kinde cool//
	// for {
	// 	fmt.Println("0")
	// }

}
