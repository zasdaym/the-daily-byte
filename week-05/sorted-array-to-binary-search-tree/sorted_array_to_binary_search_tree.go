package sortedarraytobinarysearchtree

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func NewBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Value: nums[mid],
		Left:  NewBST(nums[:mid]),
		Right: NewBST(nums[mid+1:]),
	}
	return root
}
