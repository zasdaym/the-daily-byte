package movingaverage

import "container/list"

type MovingAverage struct {
	size int
	nums *list.List
}

func New(size int) *MovingAverage {
	return &MovingAverage{
		size: size,
		nums: list.New(),
	}
}

func (ma *MovingAverage) Next(n int) int {
	ma.nums.PushBack(n)
	if ma.nums.Len() > ma.size {
		ma.nums.Remove(ma.nums.Front())
	}
	var sum int
	for e := ma.nums.Front(); e != nil; e = e.Next() {
		sum += e.Value.(int)
	}
	return sum / ma.nums.Len()
}
