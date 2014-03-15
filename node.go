package gorg

import (
	"fmt"
)

// text can be comprised of multiple lines
// position is the headline's asterisk count
type Node struct {
	headline string
	position int
	section  []string
	parent   *Node
}

func (self *Node) findParent(tree Tree) *Node {
	if tree.isEmpty() {
		return nil
	} else if tree.nodes[len(tree.nodes)-1].position < self.position {
		return tree.nodes[len(tree.nodes)-1]
	} else {
		tree.nodes = tree.nodes[0 : len(tree.nodes)-1]
		return self.findParent(tree)
	}
}

// the headline gets an <h?> tag, with ? determined by the position
// each line of text is a paragraph within a level div
func (self Node) toHtml() string {
	position := self.position
	if position == 0 {
		position = 1
	}

	var body string
	if len(self.section) > 0 {
		var text string
		for _, line := range self.section {
			text = fmt.Sprintf("%s<p>%s</p>", text, line)
		}

		body = fmt.Sprintf(
			"<div class=\"level-%d\">%s</div>",
			position,
			text,
		)
	}

	return fmt.Sprintf(
		"<h%d>%s</h%d>%s",
		position,
		self.headline,
		position,
		body,
	)
}
