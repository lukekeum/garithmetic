package main

import (
	"os"

	"github.com/lukekeum/garithmetic/compiler"
)

func main() {
	fileName := os.Args

	if len(os.Args) <= 1 {
		panic("file name argument required")
	}

	isError := compiler.Execute(fileName[1])

	if isError != 0 {
		panic("Something wrong")
	}
}
