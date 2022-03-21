package returnstartofcycle

import "github.com/zasdaym/the-daily-byte/linkedlist"

func StartOfCycle(head *linkedlist.Node[int]) *linkedlist.Node[int] {
	fast, slow := &linkedlist.Node[int]{Next: head}, &linkedlist.Node[int]{Next: head}
	for fast != nil && slow != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}
	return fast
}
