package lowestcommonancestor

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func LowestCommonAncestor(root, a, b *TreeNode) *TreeNode {
	if root == nil || root.Value == a.Value || root.Value == b.Value {
		return root
	}
	left := LowestCommonAncestor(root.Left, a, b)
	right := LowestCommonAncestor(root.Right, a, b)
	if left == nil && right == nil {
		return nil
	}
	return root
}
