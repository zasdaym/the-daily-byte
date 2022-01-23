package reverselist

import "github.com/zasdaym/the-daily-byte/linkedlist"

func Reverse(head *linkedlist.Node) *linkedlist.Node {
	var prev *linkedlist.Node
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}
