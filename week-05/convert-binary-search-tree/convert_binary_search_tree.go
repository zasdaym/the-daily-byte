package convertbinarysearchtree

import (
	"github.com/zasdaym/the-daily-byte/linkedlist"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func Convert(root *TreeNode) *linkedlist.Node {
	values := getValues(root)
	return linkedlist.New(values)
}

func getValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	if root.Left != nil {
		result = append(result, getValues(root.Left)...)
	}
	result = append(result, root.Value)
	if root.Right != nil {
		result = append(result, getValues(root.Right)...)
	}
	return result
}
