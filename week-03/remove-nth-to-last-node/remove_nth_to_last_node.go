package removenthtolastnode

import "github.com/zasdaym/the-daily-byte/linkedlist"

func RemoveNthLastNode(head *linkedlist.Node, n int) *linkedlist.Node {
	dummyHead := linkedlist.Node{
		Next: head,
	}
	back, front := &dummyHead, &dummyHead
	var i int
	for front.Next != nil && i < n {
		front = front.Next
		i++
	}
	if i < n {
		return head
	}
	for front.Next != nil {
		back, front = back.Next, front.Next
	}
	back.Next = back.Next.Next
	return dummyHead.Next
}
