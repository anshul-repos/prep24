package main

import "fmt"

// Stack structure
type Stack struct {
	items []int
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Push adds an item to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Pop from an empty stack")
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

// Peek returns the top item from the stack without removing it
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		panic("Peek from an empty stack")
	}
	return s.items[len(s.items)-1]
}

// SortStack recursively sorts the stack
func SortStack(stack *Stack) {
	if !stack.IsEmpty() {
		// Remove the top element
		top := stack.Pop()

		// Recursively sort the remaining stack
		SortStack(stack)

		// Insert the top element back in sorted order
		insertInSortedOrder(stack, top)
	}
}

// Insert an element into a sorted stack
func insertInSortedOrder(stack *Stack, element int) {
	if stack.IsEmpty() || element > stack.Peek() {
		stack.Push(element)
		return
	}

	// Remove the top element
	top := stack.Pop()

	// Recursively insert the element in the sorted stack
	insertInSortedOrder(stack, element)

	// Push the top element back
	stack.Push(top)
}

func main() {

	stack := &Stack{}
	stack.Push(3)
	stack.Push(1)
	stack.Push(4)
	stack.Push(2)
	stack.Push(5)

	fmt.Println("Original stack:", stack.items)

	SortStack(stack)

	fmt.Println("Sorted stack:", stack.items)
}
