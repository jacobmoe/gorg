package gorg

import (
	"fmt"
)

type Level struct {
	headline string
	position int
	text     string
}

func (self Level) toHtml() string {
	position := self.position
	if position == 0 {
		position = 1
	}

	paragraph := ""
	if self.text != "" {
		paragraph = fmt.Sprintf(
			"<p class=\"level-%d\">%s</p>",
			position,
			self.text,
		)
	}

	return fmt.Sprintf(
		"<h%d>%s</h%d>%s",
		position,
		self.headline,
		position,
		paragraph,
	)
}
