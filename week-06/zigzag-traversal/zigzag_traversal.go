package zigzagtraversal

import "container/list"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

type queueItem struct {
	node  *TreeNode
	level int
}

func ZigzagTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		result    [][]int
		lastLevel int
		row       []int
	)
	queue := list.New()
	queue.PushBack(queueItem{
		node:  root,
		level: 0,
	})
	for queue.Len() != 0 {
		curr := queue.Remove(queue.Front()).(queueItem)
		if lastLevel != curr.level {
			if lastLevel%2 == 1 {
				reverse(row)
			}
			result = append(result, row)
			row = []int{}
			lastLevel = curr.level
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
	return result
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
