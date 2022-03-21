package removevalue

import "github.com/zasdaym/the-daily-byte/linkedlist"

func Remove(head *linkedlist.Node[int], n int) *linkedlist.Node[int] {
	dummyHead := linkedlist.Node[int]{Next: head}
	curr := &dummyHead
	for curr.Next != nil {
		if curr.Next.Value == n {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return dummyHead.Next
}
