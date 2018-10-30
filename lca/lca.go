// https://github.com/berryjam/leetcode-golang/blob/master/lowest-common-ancestor-of-a-binary-tree.go

package lca

// finds lowest common ancestor of two given nodes and returns the LCA as a TreeNode
func LowestCommonAncestor(p, q *Node) *Node {
	pVector := make(map[int][]Node)
	distance := 0
	pVector[distance] = append(pVector[distance], p)



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

func getDistanceVector(a *Node) map[rune][]*Node {
	aVector := make(map[int][]*Node)
	distance := 0
	aVector[distance] = append(aVector[distance], a)

	queue := make([]*Node, 0, len(a.Parents))
	queue = append(queue, a)

	for i := 0; i < len(queue); i++ {
		node := queue[i]



		node.seen = true
		for _, p := range node.Parents {
			if !p.seen {
				p.seen = true
				queue = append(queue, p)
			}
		}
	}
	finishList := make([]*Node, 0, len(queue))
	for i := range queue {
		finishList = append(finishList, queue[i])
	}
}

// returns true if ancestor node is a valid ancestor of descendant node
func IsAncestor(ancestor, descendant *Node) bool {
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
