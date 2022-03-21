package mergelinkedlists

import "github.com/zasdaym/the-daily-byte/linkedlist"

// Merge merges two sorted linked lists in ascending order.
func Merge(a, b *linkedlist.Node[int]) *linkedlist.Node[int] {
	var dummyHead linkedlist.Node[int]
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
