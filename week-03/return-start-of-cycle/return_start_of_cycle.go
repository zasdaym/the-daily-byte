package returnstartofcycle

import "github.com/zasdaym/the-daily-byte/linkedlist"

func StartOfCycle(head *linkedlist.Node) *linkedlist.Node {
	fast, slow := &linkedlist.Node{Next: head}, &linkedlist.Node{Next: head}
	for fast != nil && slow != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}
	return fast
}
