package compiler

import (
	"fmt"
	"os"

	"github.com/lukekeum/garithmetic/compiler/lexer"
	"github.com/lukekeum/garithmetic/store"
)

func Execute(fileName string) int {

	fmt.Println("Compiler: Start Compiling")

	data, err := os.ReadFile(fileName)

	if err != nil {
		panic("[Error] Error occured while reading file")
	}

	content := string(data[:])
	store := []store.Token{}

	lexer.Execute(content, store)

	return 0
}
