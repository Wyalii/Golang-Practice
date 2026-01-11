package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello this is a file text reader")
	fmt.Println("please write name of you file:")
	var fileName string
	fmt.Scanln(&fileName)
	fmt.Println("reading a contents of a file: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	input := bufio.NewScanner(file)
	input.Scan()
	println(input.Text())
	file.Close()
}
