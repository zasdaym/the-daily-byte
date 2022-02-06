package findthemode

import (
	"container/list"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func FindMode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	frequencies := make(map[int]int)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		frequencies[node.Value]++
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	result, maxFreq := 0, 0
	for num, freq := range frequencies {
		if freq > maxFreq {
			result, maxFreq = num, freq
		}
	}
	return result
}
