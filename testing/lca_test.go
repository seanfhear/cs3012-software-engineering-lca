package main

import (
	"testing"
)

// Test a valid ancestor
func TestIsAncestor(t *testing.T) {
	expected := true
	actual := IsAncestor(args)
	if actual != expected {
		t.Errorf("Test failed, expected: '%t', got: '%t'", expected, actual)
	}
}

// Test an invalid ancestor
func TestIsNotAncestor(t *testing.T) {
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
