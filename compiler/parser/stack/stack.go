package stack

import (
	"fmt"
)

type Stack struct {
	stack []string
}

func New() *Stack {
	var stack []string

	return &Stack{stack}
}

func (s *Stack) Put(method string) {
	s.stack = append(s.stack, method)
	fmt.Printf("%s\n", method)
}

func (s *Stack) Push(value string) {
	s.stack = append(s.stack, value)
	fmt.Printf("push %s\n", value)
}
