package tests

import (
	"testing"
	"../lca"
)

// TODO add code coverage testing

// Test a valid ancestor
func TestBinaryIsAncestor(t *testing.T) {
	testTree := lca.MakeBinaryTree()
	expected := true
	actual := lca.IsAncestor(testTree['a'], testTree['i'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test an invalid ancestor
func TestBinaryIsNotAncestor(t *testing.T) {
	testTree := lca.MakeBinaryTree()
	expected := false
	actual := lca.IsAncestor(testTree['c'], testTree['e'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test where node a is a descendant of node b
func TestBinaryInvertedAncestor(t *testing.T) {
	testTree := lca.MakeBinaryTree()
	expected := false
	actual := lca.IsAncestor(testTree['i'], testTree['a'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test a valid LCA is returned
func TestBinaryValidLCA(t *testing.T) {
	testTree := lca.MakeBinaryTree()
	expected := testTree['b'].Val
	actual := lca.LowestCommonAncestor(testTree['a'], testTree['d'], testTree['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when one of the nodes is the LCA
func TestBinaryOneNodeIsLCA(t *testing.T) {
	testTree := lca.MakeBinaryTree()
	expected := testTree['b'].Val
	actual := lca.LowestCommonAncestor(testTree['a'], testTree['b'], testTree['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when both given nodes are the same
func TestBinaryDuplicateNode(t *testing.T) {
	testTree := lca.MakeBinaryTree()
	expected := testTree['e'].Val
	actual := lca.LowestCommonAncestor(testTree['a'], testTree['e'], testTree['e']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test a valid ancestor
func TestDAGIsAncestor(t *testing.T) {
	testDAG := lca.MakeDAG()
	expected := true
	actual := lca.IsAncestor(testDAG['a'], testDAG['i'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test an invalid ancestor
func TestDAGIsNotAncestor(t *testing.T) {
	testDAG := lca.MakeDAG()
	expected := false
	actual := lca.IsAncestor(testDAG['c'], testDAG['e'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test where node a is a descendant of node b
func TestDAGInvertedAncestor(t *testing.T) {
	testDAG := lca.MakeDAG()
	expected := false
	actual := lca.IsAncestor(testDAG['i'], testDAG['a'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test a valid LCA is returned
func TestDAGValidLCA(t *testing.T) {
	testDAG := lca.MakeDAG()
	expected := testDAG['b'].Val
	actual := lca.LowestCommonAncestor(testDAG['a'], testDAG['d'], testDAG['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when one of the nodes is the LCA
func TestDAGOneNodeIsLCA(t *testing.T) {
	testDAG := lca.MakeDAG()
	expected := testDAG['b'].Val
	actual := lca.LowestCommonAncestor(testDAG['a'], testDAG['b'], testDAG['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when both given nodes are the same
func TestDAGDuplicateNode(t *testing.T) {
	testDAG := lca.MakeDAG()
	expected := testDAG['e'].Val
	actual := lca.LowestCommonAncestor(testDAG['a'], testDAG['e'], testDAG['e']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}
