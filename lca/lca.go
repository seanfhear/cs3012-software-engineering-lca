// https://github.com/berryjam/leetcode-golang/blob/master/lowest-common-ancestor-of-a-binary-tree.go

package lca

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

func IsAncestor(curNode, p *TreeNode) bool {
    if curNode == p { // we allow a node to be a descendant of itself
        return true
    }

    leftRes := false
    if curNode.Left != nil {
        if curNode.Left == p {
            leftRes = true
        } else {
            leftRes = IsAncestor(curNode.Left, p)
        }
    }

    rightRes := false
    if curNode.Right != nil {
        if curNode.Right == p {
            rightRes = true
        } else {
            rightRes = IsAncestor(curNode.Right, p)
        }
    }

    return leftRes || rightRes
}