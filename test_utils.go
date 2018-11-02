package lca

// returns a sample binary tree using the graph library
func MakeBinaryTree() *Graph {
	g := New()
	for r := 'a'; r <= 'i'; r++ {
		g.AddVertex(r, r)
	}
	g.AddEdge('a', 'b')
	g.AddEdge('a', 'c')
	g.AddEdge('b', 'd')
	g.AddEdge('b', 'e')
	g.AddEdge('e', 'g')
	g.AddEdge('e', 'h')
	g.AddEdge('c', 'f')
	g.AddEdge('f', 'i')

	return g
}

// returns a sample DAG used for testing
func MakeDAG() *Graph {
	g := New()
	for r := 'a'; r <= 'g'; r++ {
		g.AddVertex(r, r)
	}
	g.AddEdge('a', 'b')
	g.AddEdge('b', 'c')
	g.AddEdge('b', 'd')
	g.AddEdge('b', 'e')
	g.AddEdge('c', 'e')
	g.AddEdge('d', 'e')
	g.AddEdge('e', 'f')
	g.AddEdge('g', 'd')

	return g
}
