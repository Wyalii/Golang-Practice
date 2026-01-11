package main

import (
	"fmt"
	"os"
)

func usingJoinForCommandLineArgument() {
	for i, arg := range os.Args[1:] {
		fmt.Println("index:", i, "element:", arg)
	}
}
