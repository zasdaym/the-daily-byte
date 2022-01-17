package mergelinkedlists

import "fmt"

// Node represents a linked list node.
type Node struct {
	Value int
	Next  *Node
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

// Merge merges two sorted linked lists in ascending order.
func Merge(a, b *Node) *Node {
	var dummyHead Node
	tail := &dummyHead
	for a != nil && b != nil {
		if a.Value < b.Value {
			tail.Next = a
			tail = tail.Next
			a = a.Next
		} else {
			tail.Next = b
			tail = tail.Next
			b = b.Next
		}
	}
	if a != nil {
		tail.Next = a
	}
	if b != nil {
		tail.Next = b
	}
	return dummyHead.Next
}
