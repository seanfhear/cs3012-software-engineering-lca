package lca

// Finds lowest common ancestor of two given nodes and returns the LCA as a Node.
// Returns nil if the two nodes have no common ancestor
func GetLowestCommonAncestor(g *Graph, p, q *Node) *Node {
	pVector := make(map[*Node]int)
	qVector := make(map[*Node]int)
	getDistanceVectorInit(g, p, pVector)
	getDistanceVectorInit(g, q, qVector)

	first := true
	lowestDistance := 0
	var lowestCommonAncestor *Node = nil

	for k, v := range pVector {
		if _, ok := qVector[k]; ok {
			distance := v + qVector[k]
			if first {
				lowestCommonAncestor = k
				lowestDistance = distance
				first = false
			}
			if distance < lowestDistance {
				lowestCommonAncestor = k
				lowestDistance = distance
			}
		}
	}

	return lowestCommonAncestor
}

// Helper function to reset all nodes back to unseen after GDV
func getDistanceVectorInit(g *Graph, a *Node, aVector map[*Node]int) {
	getDistanceVector(0, a, aVector)
	for _, node := range g.Nodes {
		node.Seen = false
	}
}

// Depth-first search through the graph from the source node to all ancestors
// recording distance of each ancestor from the source
func getDistanceVector(distance int, a *Node, aVector map[*Node]int) {
	a.Seen = true
	for _, p := range a.Parents {
		if !p.Seen {
			getDistanceVector(distance+1, p, aVector)
		}
	}
	aVector[a] = distance
}
