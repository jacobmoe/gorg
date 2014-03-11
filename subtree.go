package gorg

import (
	"fmt"
)

type Subtree struct {
	levels   []*Level
	branches []*Subtree
	parent   *Subtree
}

func (self Subtree) root() *Subtree {
	if self.parent == nil {
		return &self
	}

	return self.parent.root()
}

func (self *Subtree) addLevel(level *Level) {
	self.levels = append(self.levels, level)
}

func (self *Subtree) addBranch(subtree *Subtree) {
	subtree.parent = self
	self.branches = append(self.branches, subtree)
}

func (self Subtree) isEmpty() bool {
	return len(self.levels) == 0
}

func (self *Subtree) lastLevel() *Level {
	if self.isEmpty() {
		return nil
	}

	return self.levels[len(self.levels)-1]
}

func (self Subtree) toHtml() string {
	var html = "<div class=\"subtrees\">"
	html = self.branchesToHtml(html)

	return fmt.Sprintf("%s%s", html, "</div>")
}

func (self Subtree) branchesToHtml(html string) string {
	for _, level := range self.levels {
		html = fmt.Sprintf("%s%s", html, level.toHtml())
	}

	for _, branch := range self.branches {
		html = branch.branchesToHtml(html)
	}

	return html
}
