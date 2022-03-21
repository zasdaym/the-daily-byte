package findmiddleelement

import "github.com/zasdaym/the-daily-byte/linkedlist"

func Middle(head *linkedlist.Node[int]) int {
	front := &linkedlist.Node[int]{Next: head}
	back := front
	for front != nil {
		front, back = front.Next, back.Next
		if front != nil {
			front = front.Next
		}
	}
	return back.Value
}
