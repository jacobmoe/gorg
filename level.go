package gorg

import (
	"fmt"
)

// a level represents an org-mode headline, including following text
// text can be comprised of multiple lines
// position is the headline's asterisk count
type Level struct {
	headline string
	position int
	text     []string
}

// the headline gets an <h?> tag, with ? determined by the position
// each line of text is a paragraph within a level div
func (self Level) toHtml() string {
	position := self.position
	if position == 0 {
		position = 1
	}

	var body string
	if len(self.text) > 0 {
		var text string
		for _, line := range self.text {
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
