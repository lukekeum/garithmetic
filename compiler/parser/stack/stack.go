package stack

import (
	"fmt"
)

type Stack struct {
	stack  []string
	stacks [][]string
}

func New() *Stack {
	var stack []string
	var stacks [][]string

	return &Stack{stack, stacks}
}

func (s *Stack) Fin() {
	stack := s.stack[0:]
	s.stack = []string{}
	s.stacks = append(s.stacks, stack)
	fmt.Printf("semi\n")
}

func (s *Stack) Put(method string) {
	s.stack = append(s.stack, method)
	fmt.Printf("%s\n", method)
}

func (s *Stack) Push(value string) {
	s.stack = append(s.stack, value)
	fmt.Printf("push %s\n", value)
}
