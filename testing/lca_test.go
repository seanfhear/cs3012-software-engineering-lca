package testing

import (
	"testing"
	"../lca"
)

// returns a sample tree used for testing
func makeTree() map[rune]*lca.TreeNode {
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

// Test a valid ancestor
func TestIsAncestor(t *testing.T) {
	testTree := makeTree()
	expected := true
	actual := lca.IsAncestor(testTree['a'], testTree['i'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test an invalid ancestor
func TestIsNotAncestor(t *testing.T) {
	testTree := makeTree()
	expected := false
	actual := lca.IsAncestor(testTree['c'], testTree['e'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test a valid LCA is returned
func TestValidLCA(t *testing.T) {

}

// Test when one of the nodes is the LCA
func TestOneNodeIsLCA(t *testing.T) {

}

// Test when both given nodes are the same
func TestDuplicateNode(t *testing.T) {

}

// Test when wrong amount of nodes given
func TestIncorrectNumParameters(t *testing.T) {

}

// Test invalid parameters
func TestInvalidParameters(t *testing.T) {

}
