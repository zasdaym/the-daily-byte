package visiblevalues

import "container/list"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

type queueItem struct {
	node  *TreeNode
	level int
}

func VisibleValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(queueItem{
		node:  root,
		level: 0,
	})
	lastLevel := -1
	var result []int
	for queue.Len() != 0 {
		curr := queue.Remove(queue.Front()).(queueItem)
		if curr.level != lastLevel {
			lastLevel = curr.level
			result = append(result, curr.node.Value)
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
	return result
}
