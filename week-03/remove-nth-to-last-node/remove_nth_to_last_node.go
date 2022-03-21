package removenthtolastnode

import "github.com/zasdaym/the-daily-byte/linkedlist"

func RemoveNthLastNode(head *linkedlist.Node[int], n int) *linkedlist.Node[int] {
	dummyHead := linkedlist.Node[int]{
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
