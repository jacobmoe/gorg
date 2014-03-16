package gorg

import "fmt"

type Tree struct {
	nodes []*Node
}

func (self *Tree) addNode(node *Node) {
	node.parent = node.findParent(*self)
	self.nodes = append(self.nodes, node)
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
