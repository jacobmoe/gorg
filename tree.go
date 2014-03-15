package gorg

type Tree struct {
	nodes []*Node
}

func (self Tree) addNode(node *Node) {
	node.parent = node.findParent(self)
	self.nodes = append(self.nodes, node)
}

func (self Tree) isEmpty() bool {
	return len(self.nodes) == 0
}
