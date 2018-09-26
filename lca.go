// https://github.com/berryjam/leetcode-golang/blob/master/lowest-common-ancestor-of-a-binary-tree.go

package main

import (
    ds "github.com/berryjam/leetcode-golang/datastructure"
    "fmt"
)

func LowestCommonAncestor(root, p, q *ds.TreeNode) *ds.TreeNode {
    if root.Left == nil && root.Right == nil {
        if isAncestor(root, p) && isAncestor(root, q) {
            return root
        } else {
            return nil
        }
    }

    if root.Left != nil {
        if isAncestor(root.Left, p) && isAncestor(root.Left, q) {
            return LowestCommonAncestor(root.Left, p, q)
        }
    }

    if root.Right != nil {
        if isAncestor(root.Right, p) && isAncestor(root.Right, q) {
            return LowestCommonAncestor(root.Right, p, q)
        }
    }

    if isAncestor(root, p) && isAncestor(root, q) {
        return root
    }

    return nil
}

func isAncestor(curNode, p *ds.TreeNode) bool {
    if curNode == p { // we allow a node to be a descendant of itself
        return true
    }

    leftRes := false
    if curNode.Left != nil {
        if curNode.Left == p {
            leftRes = true
        } else {
            leftRes = isAncestor(curNode.Left, p)
        }
    }

    rightRes := false
    if curNode.Right != nil {
        if curNode.Right == p {
            rightRes = true
        } else {
            rightRes = isAncestor(curNode.Right, p)
        }
    }

    return leftRes || rightRes
}

func TestLowestCommonAncestor(root, p, q *ds.TreeNode) {
    res := LowestCommonAncestor(root, p, q)
    fmt.Printf("Given root=%d p=%d q=%d,LCA=%d\n", root.Val, p.Val, q.Val, res.Val)
}

func main() {
    threeNode := ds.NewTreeNode(3)
    fiveNode := ds.NewTreeNode(5)
    oneNode := ds.NewTreeNode(1)
    sixNode := ds.NewTreeNode(6)
    twoNode := ds.NewTreeNode(2)
    zeroNode := ds.NewTreeNode(0)
    eightNode := ds.NewTreeNode(8)
    sevenNode := ds.NewTreeNode(7)
    fourNode := ds.NewTreeNode(4)

    threeNode.Left = fiveNode
    threeNode.Right = oneNode
    fiveNode.Left = sixNode
    fiveNode.Right = twoNode
    oneNode.Left = zeroNode
    oneNode.Right = eightNode
    twoNode.Left = sevenNode
    twoNode.Right = fourNode

    TestLowestCommonAncestor(threeNode, sevenNode, zeroNode)
    TestLowestCommonAncestor(threeNode, sevenNode, fourNode)
    TestLowestCommonAncestor(threeNode, sixNode, fourNode)
}