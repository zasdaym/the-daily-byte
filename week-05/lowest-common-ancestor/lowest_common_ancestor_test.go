package lowestcommonancestor

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	t.Run("LCA is root", func(t *testing.T) {
		a := &TreeNode{Value: 1}
		b := &TreeNode{Value: 9}
		root := &TreeNode{
			Value: 7,
			Left: &TreeNode{
				Value: 2,
				Left:  a,
				Right: &TreeNode{Value: 5},
			},
			Right: b,
		}
		want := root

		if got := LowestCommonAncestor(root, a, b); got != want {
			t.Errorf("LowestCommonAncestor(%v, %d, %d) = %d, want %d", root, a.Value, b.Value, got.Value, want.Value)
		}
	})

	t.Run("LCA in subtree", func(t *testing.T) {
		a := &TreeNode{Value: 2}
		b := &TreeNode{Value: 6}
		root := &TreeNode{
			Value: 8,
			Left: &TreeNode{
				Value: 3,
				Left:  a,
				Right: b,
			},
			Right: &TreeNode{Value: 9},
		}
		want := root

		if got := LowestCommonAncestor(root, a, b); got != want {
			t.Errorf("LowestCommonAncestor(%v, %d, %d) = %d, want %d", root, a.Value, b.Value, got.Value, want.Value)
		}
	})
}
