package containscycle

import "github.com/zasdaym/the-daily-byte/linkedlist"

func ContainsCycle(head *linkedlist.Node) bool {
	slow, fast := &linkedlist.Node{Next: head}, &linkedlist.Node{Next: head}
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
