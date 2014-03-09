package gorg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	in  Level
	out string
}{
	{
		in: Level{
			headline: "the title",
			position: 1,
			text:     "the text",
		},
		out: "<h1>the title</h1><p class=\"level-1\">the text</p>",
	},
	{
		in: Level{
			headline: "the title2",
			position: 2,
		},
		out: "<h2>the title2</h2>",
	},
	{
		in: Level{
			headline: "the title3",
		},
		out: "<h1>the title3</h1>",
	},
}

func TestLevelToHtml(t *testing.T) {
	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}
}
