package stack

import (
	"fmt"
	"strconv"
)

type Stack struct {
	stack  []string
	result int
}

func New() *Stack {
	var stack []string
	result := 0

	return &Stack{stack, result}
}

func (s *Stack) Put(method string) {
	s.stack = append(s.stack, method)
}

func (s *Stack) Push(value int) {
	s.stack = append(s.stack, strconv.Itoa(value))
	fmt.Printf("push %d", value)
}
