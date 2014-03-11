package gorg

import (
	"fmt"
)

type Subtree struct {
	levels []Level
}

func (self Subtree) addLevel(level Level) Subtree {
	return Subtree{levels: append(self.levels, level)}
}

func (self Subtree) isEmpty() bool {
	return len(self.levels) == 0
}

func (self *Subtree) lastLevel() Level {
	if self.isEmpty() {
		return Level{}
	}

	return self.levels[len(self.levels)-1]
}

func (self Subtree) toHtml() string {
	var html = "<div class=\"subtree\">"
	for _, level := range self.levels {
		html = fmt.Sprintf("%s%s", html, level.toHtml())
	}

	return fmt.Sprintf("%s</div>", html)
}
