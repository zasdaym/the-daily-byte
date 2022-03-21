// Package linkedlist provides linked list implementation.
package linkedlist

import "fmt"

// Node represents a linked list node.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// New creates a new linked list node from given values.
func New[T comparable](values []T) *Node[T] {
	var dummyHead Node[T]
	curr := &dummyHead
	for _, v := range values {
		curr.Next = &Node[T]{Value: v}
		curr = curr.Next
	}
	return dummyHead.Next
}

// String returns string representation of a linked list.
func (n *Node[T]) String() string {
	var result string
	for n != nil {
		result = fmt.Sprintf("%s->%v", result, n.Value)
		n = n.Next
	}
	return result
}

// Equal checks if n is equal with other.
func (n *Node[T]) Equal(other *Node[T]) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if n.Value != other.Value {
		return false
	}
	return n.Next.Equal(other.Next)
}
