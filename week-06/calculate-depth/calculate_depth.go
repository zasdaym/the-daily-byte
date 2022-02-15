package calculatedepth

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CalculateDepth(root *TreeNode) int {
	return calculateDepth(root, 0)
}

func calculateDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	var (
		left  = depth + 1
		right = depth + 1
	)
	if root.Left != nil {
		left = calculateDepth(root.Left, depth+1)
	}
	if root.Right != nil {
		right = calculateDepth(root.Right, depth+1)
	}
	return max(left, right)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
