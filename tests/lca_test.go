package tests

import (
	"../lca"
	"testing"
)

// TODO add code coverage testing

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
	actual := lca.GetLowestCommonAncestor(testTree, testTree.Nodes['a'], testTree.Nodes['b']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when both given nodes are the same
func TestBinaryDuplicateNode(t *testing.T) {
	testTree := MakeBinaryTree()
	expected := testTree.Nodes['e'].Val
	actual := lca.GetLowestCommonAncestor(testTree, testTree.Nodes['a'], testTree.Nodes['e']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

/*

// Test a valid LCA is returned
func TestDAGValidLCA(t *testing.T) {
	testDAG := MakeDAG()
	expected := testDAG['b'].Val
	actual := lca.LowestCommonAncestor(testDAG['a'], testDAG['d'], testDAG['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when one of the nodes is the LCA
func TestDAGOneNodeIsLCA(t *testing.T) {
	testDAG := MakeDAG()
	expected := testDAG['b'].Val
	actual := lca.LowestCommonAncestor(testDAG['a'], testDAG['b'], testDAG['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when both given nodes are the same
func TestDAGDuplicateNode(t *testing.T) {
	testDAG := MakeDAG()
	expected := testDAG['e'].Val
	actual := lca.LowestCommonAncestor(testDAG['a'], testDAG['e'], testDAG['e']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}
*/