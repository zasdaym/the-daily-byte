package gathernarytreelevels

import "container/list"

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

func GatherNaryTreeLevels(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		lastLevel int
		row       []int
		result    [][]int
	)
	type queueItem struct {
		node  *TreeNode
		level int
	}
	queue := list.New()
	queue.PushBack(queueItem{
		node:  root,
		level: lastLevel,
	})
	for queue.Len() != 0 {
		curr := queue.Remove(queue.Front()).(queueItem)
		if curr.level != lastLevel {
			result = append(result, row)
			row = []int{}
			lastLevel = curr.level
		}
		row = append(row, curr.node.Value)
		for _, child := range curr.node.Children {
			queue.PushBack(queueItem{
				node:  child,
				level: curr.level + 1,
			})
		}
	}
	result = append(result, row)
	return result
}
