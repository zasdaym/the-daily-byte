package maxvalueineachlevel

import "container/list"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

type queueItem struct {
	node  *TreeNode
	level int
}

func MaxValuePerLevel(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(queueItem{
		node:  root,
		level: 0,
	})
	var (
		result    []int
		max       int
		prevLevel int
	)
	for queue.Len() != 0 {
		curr := queue.Remove(queue.Front()).(queueItem)
		if curr.level != prevLevel {
			result = append(result, max)
			max = 0
			prevLevel = curr.level
		}
		if curr.node.Value > max {
			max = curr.node.Value
		}
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
	result = append(result, max)
	return result
}
