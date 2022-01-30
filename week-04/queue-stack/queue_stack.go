package queuestack

import "container/list"

// Stack represents stack implemented with single queue.
type Stack struct {
	queue *list.List
}

// New creates a new Stack.
func New(values []int) *Stack {
	queue := list.New()
	for _, v := range values {
		queue.PushBack(v)
	}
	return &Stack{
		queue: queue,
	}
}

// Push pushes given value into top of the stack.
func (s *Stack) Push(value int) {
	s.queue.PushBack(value)
}

// Pop removes the top element of the stack.
func (s *Stack) Pop() {
	s.queue.Remove(s.queue.Back())
}

// Peek returns the top element of the stack.
func (s *Stack) Peek() int {
	return s.queue.Back().Value.(int)
}

// Empty returns whether or not the stack is empty.
func (s *Stack) Empty() bool {
	return s.queue.Len() == 0
}
