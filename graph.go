// based on https://github.com/funkygao/golib/blob/master/dag/dag.go

package lca

type Graph struct {
	Nodes map[rune]*Node
}

type Node struct {
	Name rune
	Val  rune

	Seen     bool
	Parents  []*Node
}

func New() *Graph {
	this := new(Graph)
	this.Nodes = make(map[rune]*Node)
	return this
}

func (this *Graph) AddVertex(name rune, val rune) *Node {
	node := &Node{Name: name, Val: val}
	this.Nodes[name] = node
	return node
}

func (this *Graph) AddEdge(from, to rune) {
	fromNode := this.Nodes[from]
	toNode := this.Nodes[to]
	toNode.Parents = append(toNode.Parents, fromNode)
}
