// https://github.com/berryjam/leetcode-golang/blob/master/datastructure/tree-node.go

package lca

type TreeNode struct {
	Val   rune
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val rune) *TreeNode {
	return &TreeNode{
		Val:val,
	}
}