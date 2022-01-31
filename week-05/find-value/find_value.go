package findvalue

type Node struct {
	Value       int
	Left, Right *Node
}

func FindValue(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if value > root.Value {
		return FindValue(root.Right, value)
	}
	if value < root.Value {
		return FindValue(root.Left, value)
	}
	return root
}
