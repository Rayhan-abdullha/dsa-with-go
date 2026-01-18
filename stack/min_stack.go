package stack

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}
func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)

	// minStack empty হলে বা নতুন value ছোট হলে
	if len(s.minStack) == 0 || x < s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, x)
	} else {
		// আগের minimum আবার push করি
		s.minStack = append(s.minStack, s.minStack[len(s.minStack)-1])
	}
}
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}
func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func MinSolution() {
	stack := NewMinStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(0)
	stack.GetMin() // return 0
	stack.Pop()
	stack.Top()    // return 2
	stack.GetMin() // return 1
	fmt.Println(stack)
}
