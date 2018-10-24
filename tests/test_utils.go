package tests

import (
	"../lca"
	)

// returns a sample tree used for testing
func MakeBinaryTree() map[rune]*lca.TreeNode {
	nodes := make(map[rune]*lca.TreeNode)
	for r := 'a'; r <= 'i'; r++ {
		nodes[r] = lca.NewTreeNode(r)
	}
	nodes['a'].Left = nodes['b']
	nodes['a'].Right = nodes['c']
	nodes['b'].Left = nodes['d']
	nodes['b'].Right = nodes['e']
	nodes['e'].Left = nodes['g']
	nodes['e'].Right = nodes['h']
	nodes['c'].Right = nodes['f']
	nodes['f'].Right = nodes['i']
	return nodes
}


// returns a sample DAG used for testing
func MakeDAG() {

}