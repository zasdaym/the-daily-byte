// Package tree provides various tree implementation.
package tree

type BinaryTreeNode[T comparable] struct {
	Value       T
	Left, Right *BinaryTreeNode[T]
}
