package bottomsup

import "container/list"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

type queueItem struct {
	node  *TreeNode
	level int
}

func BottomsUp(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		lastLevel int
		result    [][]int
		row       []int
	)
	queue := list.New()
	queue.PushBack(queueItem{
		node:  root,
		level: lastLevel,
	})
	for queue.Len() != 0 {
		curr := queue.Remove(queue.Front()).(queueItem)
		if curr.level != lastLevel {
			lastLevel = curr.level
			result = append(result, row)
			row = []int{}
		}
		row = append(row, curr.node.Value)
		if curr.node.Left != nil {
			queue.PushBack(queueItem{
				node:  curr.node.Left,
				level: curr.level + 1,
			})
		}
		if curr.node.Right != nil {
			queue.PushBack(queueItem{
				node:  curr.node.Right,
				level: curr.level + 1,
			})
		}
	}
	result = append(result, row)

	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}
