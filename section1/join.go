package main

import (
	"fmt"
	"os"
	"strings"
)

func usingJoinForCommandLineArgument() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}
