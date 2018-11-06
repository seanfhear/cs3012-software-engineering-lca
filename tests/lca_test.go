package tests

import (
	"../lca"
	"testing"
)

// Test a valid LCA is returned
func TestBinaryValidLCA(t *testing.T) {
	testTree := MakeBinaryTree()
	expected := testTree.Nodes['b'].Val
	actual := lca.GetLowestCommonAncestor(testTree, testTree.Nodes['d'], testTree.Nodes['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when one of the nodes is the LCA
func TestBinaryOneNodeIsLCA(t *testing.T) {
	testTree := MakeBinaryTree()
	expected := testTree.Nodes['b'].Val
	actual := lca.GetLowestCommonAncestor(testTree, testTree.Nodes['b'], testTree.Nodes['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when both given nodes are the same
func TestBinaryDuplicateNode(t *testing.T) {
	testTree := MakeBinaryTree()
	expected := testTree.Nodes['e'].Val
	actual := lca.GetLowestCommonAncestor(testTree, testTree.Nodes['e'], testTree.Nodes['e']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test a valid LCA is returned
func TestDAGValidLCA(t *testing.T) {
	testDAG := MakeDAG()
	expected := testDAG.Nodes['b'].Val
	actual := lca.GetLowestCommonAncestor(testDAG, testDAG.Nodes['d'], testDAG.Nodes['c']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when one of the nodes is the LCA
func TestDAGOneNodeIsLCA(t *testing.T) {
	testDAG := MakeDAG()
	expected := testDAG.Nodes['b'].Val
	actual := lca.GetLowestCommonAncestor(testDAG, testDAG.Nodes['b'], testDAG.Nodes['e']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when both given nodes are the same
func TestDAGDuplicateNode(t *testing.T) {
	testDAG := MakeDAG()
	expected := testDAG.Nodes['e'].Val
	actual := lca.GetLowestCommonAncestor(testDAG, testDAG.Nodes['e'], testDAG.Nodes['e']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

//Test when no common ancestor
func TestDAGNoCommonAncestor(t *testing.T) {
	testDAG := MakeDAG()
	var expected *lca.Node = nil
	actual := lca.GetLowestCommonAncestor(testDAG, testDAG.Nodes['g'], testDAG.Nodes['a'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, actual)
	}
}
