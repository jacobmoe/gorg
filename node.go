package gorg

import (
	"fmt"
)

// Node represents an org-mode headline with a following section
// a section can be comprised of multiple lines
// position is the headline's asterisk count
type Node struct {
	headline string
	position int
	section  []string
	parent   *Node
}

func (self *Node) findParent(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	} else if nodes[len(nodes)-1].position < self.position {
		return nodes[len(nodes)-1]
	} else {
		nodes = nodes[0 : len(nodes)-1]
		return self.findParent(nodes)
	}
}

// the headline gets an <h?> tag, with ? determined by the position
// each line of text is a paragraph within a level div
func (self Node) toHtml() string {
	var header string

	if self.position != 0 {
		header = fmt.Sprintf(
			"<h%d>%s</h%d>",
			self.position,
			self.headline,
			self.position,
		)
	}

	var body string
	if len(self.section) > 0 {
		var text string
		for _, line := range self.section {
			text = fmt.Sprintf("%s<p>%s</p>", text, line)
		}

		body = fmt.Sprintf(
			"<div class=\"level-%d\">%s</div>",
			self.position,
			text,
		)
	}

	return fmt.Sprintf("%s%s", header, body)
}
