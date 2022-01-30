package queuestack

import "testing"

func TestQueueStack(t *testing.T) {
	stack := New([]int{1, 2, 3, 4})
	if stack.Peek() != 4 {
		t.Errorf("stack.Peek() = %d, want %d", stack.Peek(), 4)
	}

	stack.Push(5)
	if stack.Peek() != 5 {
		t.Errorf("stack.Peek() = %d, want %d", stack.Peek(), 5)
	}

	stack.Pop()
	stack.Pop()
	if stack.Peek() != 3 {
		t.Errorf("stack.Peek() = %d, want %d", stack.Peek(), 3)
	}

	if stack.Empty() {
		t.Errorf("stack.Empty() = %v, want %v", stack.Empty(), false)
	}

	stack.Pop()
	stack.Pop()
	stack.Pop()
	if !stack.Empty() {
		t.Errorf("stack.Empty() = %v, want %v", stack.Empty(), true)
	}
}
