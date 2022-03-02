package main

import "fmt"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str) 
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1 
		element := (*s)[index] 
		*s = (*s)[:index] 
		return element, true
	}
}

func main() {
	var stack Stack 
	stack.Push("3")
	stack.Push("2")
	stack.Push("1")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y{
			fmt.Println(x)
		}
	}
}