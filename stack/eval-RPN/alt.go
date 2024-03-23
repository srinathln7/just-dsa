package main

import "strconv"

type Stack struct {
	vals []int
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.vals = append(s.vals, val)
}

func (s *Stack) Pop() int {
	var x int
	n := len(s.vals)
	if n > 0 {
		x = s.vals[n-1]
		s.vals = s.vals[:n-1]
	}
	return x
}

func (s *Stack) Top() int {
	var x int
	n := len(s.vals)
	if n > 0 {
		x = s.vals[n-1]
	}
	return x
}

func (s *Stack) Len() int {
	return len((*s).vals)
}

func evalRPNAlt(tokens []string) int {

	//
	s := newStack()

	for _, token := range tokens {

		switch {
		case token == "+" || token == "-" || token == "*" || token == "/":
			if s.Len() >= 2 {
				val2 := s.Pop()
				val1 := s.Pop()
				if token == "+" {
					s.Push(val1 + val2)
				}
				if token == "-" {
					s.Push(val1 - val2)
				}
				if token == "*" {
					s.Push(val1 * val2)
				}
				if token == "/" {
					if val2 != 0 {
						s.Push(val1 / val2)
					}

				}
			}

		default:
			val, err := strconv.Atoi(token)
			if err == nil {
				s.Push(val)
			}
		}

	}

	return s.Top()

}
