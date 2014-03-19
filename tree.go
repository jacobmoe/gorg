package gorg

import "fmt"

type Tree struct {
	nodes    []*Node
	subtrees []*Tree
	parent   *Tree
}

func (self *Tree) addNode(node *Node) {
	node.parent = node.findParent(self.nodes)
	self.nodes = append(self.nodes, node)
}

func (self *Tree) addSubtree(subtree *Tree) {
	subtree.parent = self
	self.subtrees = append(self.subtrees, subtree)
}

func (self Tree) isEmpty() bool {
	return len(self.nodes) == 0
}

func (self Tree) lastNode() *Node {
	return self.nodes[len(self.nodes)-1]
}

func (self Tree) toHtml() string {
	var html string
	for _, node := range self.nodes {
		if node.parent == nil {
			if node != self.nodes[0] {
				html = html + "</div>"
			}

			html = html + "<div class=\"tree\">"
		}

		html = fmt.Sprintf("%s%s", html, node.toHtml())

		if node == self.lastNode() {
			html = fmt.Sprintf("%s</div>", html)
		}
	}

	return html
}

func (self *Tree) indexOfNode(searchNode *Node) int {
	for i, node := range self.nodes {
		if node == searchNode {
			return i
		}
	}

	return -1
}

func (self *Tree) deleteNode(node *Node) {
	i := self.indexOfNode(node)

	if i == -1 {
		return
	}

	if i == 0 {
		self.nodes = self.nodes[1:]
	} else if i == len(self.nodes)-1 {
		self.nodes = self.nodes[:len(self.nodes)-1]
	} else {
		self.nodes = append(self.nodes[:i], self.nodes[i+1:]...)
	}
}

func (self *Tree) unflatten() {
	subtrees := getSubtrees(self.nodes)

	for _, s := range subtrees {
		self.addSubtree(s)

		for _, n := range s.nodes {
			self.deleteNode(n)
		}
	}

	for _, subtree := range self.subtrees {
		subtree.unflatten()
	}
}

func getSubtrees(ns []*Node) []*Tree {
	if len(ns) == 1 {
		return []*Tree{}
	}

	root := ns[0]
	nodes := ns[1:]

	subtree := &Tree{nodes: []*Node{root}}
	var subtrees []*Tree

	for _, node := range nodes {
		if node.position > root.position {
			subtree.addNode(node)

			if node == nodes[len(nodes)-1] {
				subtrees = append(subtrees, subtree)
			}
		} else {
			subtrees = append(subtrees, subtree)

			root = node
			subtree = &Tree{nodes: []*Node{root}}
		}
	}

	if len(subtrees) > 1 {
		return subtrees
	} else {
		return getSubtrees(nodes)
	}

}
