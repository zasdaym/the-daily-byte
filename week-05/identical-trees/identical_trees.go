package identicaltrees

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func Identical(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Value != b.Value {
		return false
	}
	return Identical(a.Left, b.Left) && Identical(a.Right, b.Right)
}
