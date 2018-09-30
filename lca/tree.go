// https://github.com/berryjam/leetcode-golang/blob/master/datastructure/tree-node.go

package lca

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:val,
	}
}