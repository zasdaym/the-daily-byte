package callcounter

type CallCounter struct {
	calls []call
}

type call struct {
	timestamp int
	expired   bool
}

const timeoutMS = 3000

func (c *CallCounter) Ping(timestamp int) int {
	c.calls = append(c.calls, call{timestamp: timestamp})
	var result int
	for i := range c.calls {
		if timestamp-c.calls[len(c.calls)-1-i].timestamp > timeoutMS {
			c.calls[len(c.calls)-1-i].expired = true
			break
		}
		result++
	}
	return result
}
