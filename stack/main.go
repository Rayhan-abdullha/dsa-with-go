package stack

import "fmt"

type ArrayStack struct {
	stack    []int
	capacity int
	top      int
}

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		stack:    make([]int, capacity),
		capacity: capacity,
		top:      -1,
	}
}

func (s *ArrayStack) Push(value int) {
	if s.top == s.capacity-1 {
		fmt.Println("Stack Overflow")
		return
	}
	s.top++
	s.stack[s.top] = value
}
func (s *ArrayStack) Peek() int {
	if s.isEmpty() {
		fmt.Println("stack is empty!")
		return -1
	}
	return s.stack[s.top]
}
func (s *ArrayStack) isEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *ArrayStack) Pop() int {
	if s.isEmpty() {
		fmt.Println("stack is empty!")
		return -1
	}
	top := s.stack[s.top]
	s.top--
	return top
}
func isValid(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (ch == ')' && top != '(') ||
				(ch == '}' && top != '{') ||
				(ch == ']' && top != '[') {
				return false
			}
		}
	}
	return len(stack) == 0
}

func Stack() {
	// stack := NewArrayStack(10)
	// stack.Push(10)
	// stack.Push(1010)
	// top3 := stack.Peek()

	// fmt.Println(top3)
	fmt.Println(isValid("[{}]"))
}
