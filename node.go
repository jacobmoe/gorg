/* 
 * A Node models an org-mode headline with a following section
 * a section can be comprised of multiple lines
 * position is the headline's asterisk count
*/ 

package gorg

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Headline string   `json:"headline"`
	Position int      `json:"position"`
	Section  []string `json:"sections"`
	parent   *Node
}

func (self *Node) findParent(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	} else if nodes[len(nodes)-1].Position < self.Position {
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

	if self.Headline != "" {
		header = fmt.Sprintf(
			"<h%d>%s</h%d>",
			self.Position,
			self.Headline,
			self.Position,
		)
	}

	var body string
	if len(self.Section) > 0 {
		var text string
		for _, line := range self.Section {
			text = fmt.Sprintf("%s<p>%s</p>", text, line)
		}

		body = fmt.Sprintf(
			"<div class=\"level-%d\">%s</div>",
			self.Position,
			text,
		)
	}

	return fmt.Sprintf("%s%s", header, body)
}

func (self Node) toJson() string {
	json, err := json.Marshal(self)
	check(err)

	return string(json)
}
