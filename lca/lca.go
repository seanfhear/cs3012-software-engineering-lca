package lca

import "fmt"

// finds lowest common ancestor of two given nodes and returns the LCA as a TreeNode
func LowestCommonAncestor(g *Graph, p, q *Node) *Node {
	pVector := make(map[int][]*Node)
	qVector := make(map[int][]*Node)
	getDistanceVectorInit(g, p, pVector)
	getDistanceVectorInit(g, q, qVector)

	fmt.Println(pVector)
	fmt.Println(qVector)

	return nil
}

// helper function to reset all nodes back to unseen after GDV
func getDistanceVectorInit(g *Graph, a *Node, aVector map[int][]*Node) {
	getDistanceVector(0, a, aVector)
	for _, node := range g.Nodes {
		node.Seen = false
	}
}

// depth-first search through the graph from the source node to all ancestors
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
