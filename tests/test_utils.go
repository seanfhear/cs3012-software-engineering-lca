package tests

import (
	"../lca"
)

// returns a sample tree used for testing
func MakeBinaryTree() map[rune]*lca.TreeNode {
	nodes := make(map[rune]*lca.TreeNode)
	for r := 'a'; r <= 'i'; r++ {
		nodes[r] = lca.NewTreeNode(r)
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


// returns a sample binary tree using the graph library
func MakeBinaryTree2() *lca.Graph {
	g := lca.New()
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

/*
// returns a sample DAG used for testing
func MakeDAG() lca.Graph {
	nodes := make([]lca.Node, 7)
	g := lca.New(lca.Directed)
	for r := 'a'; r <= 'g'; r++ {
		nodes[r] = g.MakeNode()
	}
	g.MakeEdge(nodes['a'], nodes['b'])
	g.MakeEdge(nodes['b'], nodes['c'])
	g.MakeEdge(nodes['b'], nodes['d'])
	g.MakeEdge(nodes['b'], nodes['e'])
	g.MakeEdge(nodes['c'], nodes['e'])
	g.MakeEdge(nodes['d'], nodes['e'])
	g.MakeEdge(nodes['e'], nodes['f'])
	g.MakeEdge(nodes['g'], nodes['d'])

	fmt.Println(nodes['a'])

	return g
}
*/