package gorg

import (
	"fmt"
)

type Subtree struct {
	levels   []*Level
	branches []*Subtree
	parent   *Subtree
	tree     *Tree
}

func (self Subtree) root() *Subtree {
	if self.parent == nil {
		return &self
	}

	return self.parent.root()
}

// base case:
//   top branch and level position <= first level position
// place next level:
//   self has no branches and level position > last level
// add new subtree to parent:
//   level position equal to first level of a subtree
// branch off new subtrees:
//   level greater than first level and less than or equal to last
//   2 new subtrees
func (self *Subtree) addLevel(level *Level) {

	if self.isEmpty() {
		self.levels = append(self.levels, level)
		return
	}

	if level.position <= self.root().firstLevel().position {
		self.tree.addSubtree(&Subtree{levels: []Level{level}})
		return
	}

	if level.position > self.lastLevel().position {
		self.levels = append(self.levels, level)
	} else if level.position == self.firstLevel().position {
		self.parent.addBranch(&Subtree{levels: []Level{level}})
	} else if level.position < self.firstLevel().position {
		// recursive call using parent
	} else {
		// scan through levels, find
	}
}

func (self *Subtree) addBranch(subtree *Subtree) {
	subtree.parent = self
	self.branches = append(self.branches, subtree)
}

func (self Subtree) isEmpty() bool {
	return len(self.levels) == 0
}

func (self Subtree) isBottomBranch() bool {
	return len(self.branches) == 0
}

func (self *Subree) firstLevel() *Level {
	if self.isEmpty() {
		return nil
	}

	return self.levels[0]
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
