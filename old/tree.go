package gorg

import "fmt"
import "encoding/json"

type Tree struct {
	Nodes    []*Node `json:"nodes"`
	Subtrees []*Tree `json:"subtrees"`
	parent   *Tree
}

func NewTree(nodes []*Node) *Tree {
	tree := Tree{Nodes: nodes}
	tree.unflatten()

	return &tree
}

func (self *Tree) addNode(node *Node) {
	node.parent = node.findParent(self.Nodes)

	if node.Position == 0 {
		node.Position = 1
	}

	self.Nodes = append(self.Nodes, node)
}

func (self *Tree) addSubtree(subtree *Tree) {
	subtree.parent = self
	self.Subtrees = append(self.Subtrees, subtree)
}

func (self Tree) isEmpty() bool {
	return len(self.Nodes) == 0
}

func (self Tree) lastNode() *Node {
	return self.Nodes[len(self.Nodes)-1]
}

func (self Tree) toHtml() string {
	var html string

	// if top level of tree has Nodes,
	// top is one, collapsible subtree
	if len(self.Nodes) > 1 {
		html = "<div class=\"subtree\">"
	}
	html = self.subtreesToHtml(html)

	if len(self.Nodes) > 1 {
		html = html + "</div>"
	}

	return html
}

func (self Tree) subtreesToHtml(html string) string {

	for _, node := range self.Nodes {
		html = fmt.Sprintf("%s%s", html, node.toHtml())
	}

	for _, subtree := range self.Subtrees {
		html = html + "<div class=\"subtree\">"
		html = subtree.subtreesToHtml(html)
		html = html + "</div>"
	}

	return html
}

func (self Tree) toJson() []byte {
	json, err := json.Marshal(self)
	check(err)

	return json
}

func (self *Tree) indexOfNode(searchNode *Node) int {
	for i, node := range self.Nodes {
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
		self.Nodes = self.Nodes[1:]
	} else if i == len(self.Nodes)-1 {
		self.Nodes = self.Nodes[:len(self.Nodes)-1]
	} else {
		self.Nodes = append(self.Nodes[:i], self.Nodes[i+1:]...)
	}
}

func (self *Tree) unflatten() {
	subtrees := getSubtrees(self.Nodes)

	for _, s := range subtrees {
		self.addSubtree(s)

		for _, n := range s.Nodes {
			self.deleteNode(n)
		}
	}

	for _, subtree := range self.Subtrees {
		subtree.unflatten()
	}
}

func getSubtrees(ns []*Node) []*Tree {

	if len(ns) == 1 {
		return []*Tree{}
	}

	root := ns[0]
	nodes := ns[1:]

	subtree := &Tree{Nodes: []*Node{root}}
	var subtrees []*Tree

	for _, node := range nodes {
		if node.Position > root.Position {
			subtree.addNode(node)
		} else {
			subtrees = append(subtrees, subtree)

			root = node

			subtree = &Tree{Nodes: []*Node{root}}

		}

		if node == nodes[len(nodes)-1] {
			subtrees = append(subtrees, subtree)
		}
	}

	if len(subtrees) > 1 {
		return subtrees
	} else {
		return getSubtrees(nodes)
	}
}

func printTree(tree Tree) {
	for _, node := range tree.Nodes {
		line := ""
		for i := 0; i < node.Position; i++ {
			line = line + "*"
		}

		line = line + " " + node.Headline
		fmt.Println(line)
	}

	for _, subtree := range tree.Subtrees {
		printTree(*subtree)
	}
}
