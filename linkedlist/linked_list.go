// Package linkedlist provides linked list implementation.
package linkedlist

import "fmt"

// Node represents a linked list node.
type Node struct {
	Value int
	Next  *Node
}

// New creates a new linked list from given values.
func New(values []int) *Node {
	var dummyHead Node
	curr := &dummyHead
	for _, v := range values {
		curr.Next = &Node{Value: v}
		curr = curr.Next
	}
	return dummyHead.Next
}

// String returns string representation of a linked list.
func (n *Node) String() string {
	var result string
	for n != nil {
		result = fmt.Sprintf("%s->%d", result, n.Value)
		n = n.Next
	}
	return result
}

// Equal checks if n is equal with other.
func (n *Node) Equal(other *Node) bool {
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
