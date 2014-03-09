package gorg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubtreeToHtml(t *testing.T) {
	var tests = []struct {
		in  Subtree
		out string
	}{
		{
			in: Subtree{
				levels: []Level{
					Level{
						headline: "the headline1",
						position: 1,
					},
					Level{
						headline: "the headline2",
						position: 2,
						text:     "the text2",
					},
				},
			},
			out: "<div class=\"subtree\"><h1>the headline1</h1><h2>the headline2</h2><p class=\"level-2\">the text2</p></div>",
		},
		{
			in: Subtree{
				levels: []Level{
					Level{
						headline: "headline1",
						position: 1,
						text:     "text1",
					},
					Level{
						headline: "headline2",
						position: 2,
						text:     "text2",
					},
				},
			},
			out: "<div class=\"subtree\"><h1>headline1</h1><p class=\"level-1\">text1</p><h2>headline2</h2><p class=\"level-2\">text2</p></div>",
		},
	}

	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}
}
