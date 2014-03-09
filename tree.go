package gorg

import "fmt"

type Tree struct {
	subtrees []Subtree
}

func (self *Tree) addSubtree(subtree Subtree) Tree {
	return Tree{subtrees: append(self.subtrees, subtree)}
}

func (self Tree) toHtml() string {
	var html = "<div class=\"tree\">"
	for _, subtree := range self.subtrees {
		html = fmt.Sprintf("%s%s", html, subtree.toHtml())
	}

	return fmt.Sprintf("%s</div>", html)
}
