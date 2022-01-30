package callcounter

import (
	"container/list"
)

type CallCounter struct {
	calls *list.List
}

func New() *CallCounter {
	return &CallCounter{
		calls: list.New(),
	}
}

const timeoutMS = 3000

func (c *CallCounter) Ping(timestamp int) int {
	c.calls.PushBack(timestamp)
	for timestamp-c.calls.Front().Value.(int) > timeoutMS {
		c.calls.Remove(c.calls.Front())
	}
	return c.calls.Len()
}
