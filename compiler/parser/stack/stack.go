package stack

import "fmt"

type Stack struct {
	stack  []int
	result int
}

func New() *Stack {
	var stack []int
	result := 0

	return &Stack{stack, result}
}

func (s *Stack) Add() {
	a := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	b := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	s.stack = append(s.stack, a+b)
	fmt.Println("add")
}

func (s *Stack) Sub() {
	a := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	b := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	s.stack = append(s.stack, b-a)
	fmt.Println("sub")
}

func (s *Stack) Mul() {
	a := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	b := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	s.stack = append(s.stack, a*b)
	fmt.Println("mul")
}

func (s *Stack) Div() {
	a := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	b := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	s.stack = append(s.stack, b/a)
	fmt.Println("div")
}

func (s *Stack) Push(value int) {
	s.stack = append(s.stack, value)
	fmt.Printf("push %d", value)
}
