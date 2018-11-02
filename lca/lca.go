package lca

import "fmt"

// finds lowest common ancestor of two given nodes and returns the LCA as a TreeNode
func LowestCommonAncestor(p, q *Node) *Node {
	pVector := make(map[int][]*Node)
	getDistanceVector(0, p, pVector)

	fmt.Println(pVector)


	return nil
}

// breadth-first search through the graph from the source node to all ancestors
// recording distance of each ancestor from the source
func getDistanceVector(distance int, a *Node, aVector map[int][]*Node) {
	a.Seen = true

	for _, p := range a.Parents {
		if !p.Seen {
			distance += 1
			getDistanceVector(distance, p, aVector)
			distance -= 1
		}
	}
	aVector[distance] = append(aVector[distance], a)
}
