package gorg

import (
	"fmt"
)

type Subtree struct {
	levels []Level
}

func (self Subtree) toHtml() string {
	var html = "<div class=\"subtree\">"
	for _, level := range self.levels {
		html = fmt.Sprintf("%s%s", html, level.toHtml())
	}

	return fmt.Sprintf("%s</div>", html)
}
