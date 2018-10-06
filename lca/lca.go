// https://github.com/berryjam/leetcode-golang/blob/master/lowest-common-ancestor-of-a-binary-tree.go

package lca

// finds lowest common ancestor of two given nodes and returns the LCA as a TreeNode
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root.Left == nil && root.Right == nil {
        if IsAncestor(root, p) && IsAncestor(root, q) {
            return root
        } else {
            return nil
        }
    }

    if root.Left != nil {
        if IsAncestor(root.Left, p) && IsAncestor(root.Left, q) {
            return LowestCommonAncestor(root.Left, p, q)
        }
    }

    if root.Right != nil {
        if IsAncestor(root.Right, p) && IsAncestor(root.Right, q) {
            return LowestCommonAncestor(root.Right, p, q)
        }
    }

    if IsAncestor(root, p) && IsAncestor(root, q) {
        return root
    }

    return nil
}

// returns true if ancestor node is a valid ancestor of descendant node
func IsAncestor(ancestor, descendant *TreeNode) bool {
    if ancestor == descendant { // we allow a node to be a descendant of itself
        return true
    }

    leftRes := false
    if ancestor.Left != nil {
        if ancestor.Left == descendant {
            leftRes = true
        } else {
            leftRes = IsAncestor(ancestor.Left, descendant)
        }
    }

    rightRes := false
    if ancestor.Right != nil {
        if ancestor.Right == descendant {
            rightRes = true
        } else {
            rightRes = IsAncestor(ancestor.Right, descendant)
        }
    }

    return leftRes || rightRes
}


// returns a sample tree used for testing
func MakeBinaryTree() map[rune]*TreeNode {
    nodes := make(map[rune]*TreeNode)
    for r := 'a'; r <= 'i'; r++ {
        nodes[r] = NewTreeNode(r)
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
