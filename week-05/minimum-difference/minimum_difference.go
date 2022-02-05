package minimumdifference

import "math"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func MinDifference(root *TreeNode) int {
	values := treeValues(root)
	if len(values) == 0 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	result := math.MaxInt32
	for i := range values[:len(values)-1] {
		diff := values[i+1] - values[i]
		if diff < result {
			result = diff
		}
	}
	return result
}

func treeValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	leftValues := treeValues(root.Left)
	result = append(result, leftValues...)
	result = append(result, root.Value)
	rightValues := treeValues(root.Right)
	result = append(result, rightValues...)
	return result
}
