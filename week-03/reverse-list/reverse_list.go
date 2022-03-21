package reverselist

import "github.com/zasdaym/the-daily-byte/linkedlist"

func Reverse(head *linkedlist.Node[int]) *linkedlist.Node[int] {
	var prev *linkedlist.Node[int]
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}
