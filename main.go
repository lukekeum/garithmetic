package main

import (
	"fmt"
	"os"

	"github.com/lukekeum/garithmetic/compiler"
)

func main() {
	fileName := os.Args

	if len(os.Args) <= 1 {
		panic("[Error] File name argument required")
	}

	isError := compiler.Execute(fileName[1])

	if isError != 0 {
		panic("[Error] Something wrong")
	}

	fmt.Println("Compiler: Compile succeed")
}
