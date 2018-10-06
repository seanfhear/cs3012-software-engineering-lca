package tests

import (
	"testing"
	"../lca"
)

// TODO add code coverage testing

// Test a valid ancestor
func TestIsAncestor(t *testing.T) {
	testTree := lca.MakeTree()
	expected := true
	actual := lca.IsAncestor(testTree['a'], testTree['i'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test an invalid ancestor
func TestIsNotAncestor(t *testing.T) {
	testTree := lca.MakeTree()
	expected := false
	actual := lca.IsAncestor(testTree['c'], testTree['e'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test where node a is a descendant of node b
func TestInvertedAncestor(t *testing.T) {
	testTree := lca.MakeTree()
	expected := false
	actual := lca.IsAncestor(testTree['i'], testTree['a'])
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test a valid LCA is returned
func TestValidLCA(t *testing.T) {
	testTree := lca.MakeTree()
	expected := testTree['b'].Val
	actual := lca.LowestCommonAncestor(testTree['a'], testTree['d'], testTree['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when one of the nodes is the LCA
func TestOneNodeIsLCA(t *testing.T) {
	testTree := lca.MakeTree()
	expected := testTree['b'].Val
	actual := lca.LowestCommonAncestor(testTree['a'], testTree['b'], testTree['h']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}

// Test when both given nodes are the same
func TestDuplicateNode(t *testing.T) {
	testTree := lca.MakeTree()
	expected := testTree['e'].Val
	actual := lca.LowestCommonAncestor(testTree['a'], testTree['e'], testTree['e']).Val
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}
