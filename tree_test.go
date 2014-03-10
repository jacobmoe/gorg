package gorg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeToHtml(t *testing.T) {
	var tests = []struct {
		in  Tree
		out string
	}{
		{
			in: Tree{
				subtrees: []Subtree{
					Subtree{
						levels: []Level{
							Level{
								headline: "headline1.1",
								position: 1,
							},
							Level{
								headline: "headline1.2",
								position: 2,
								text:     []string{"text1.2"},
							},
						},
					},
					Subtree{
						levels: []Level{
							Level{
								headline: "headline2.1",
								position: 1,
								text:     []string{"text2.1"},
							},
							Level{
								headline: "headline2.2",
								position: 2,
								text:     []string{"text2.2"},
							},
						},
					},
				},
			},
			out: "<div class=\"tree\"><div class=\"subtree\"><h1>headline1.1</h1><h2>headline1.2</h2><div class=\"level-2\"><p>text1.2</p></div></div><div class=\"subtree\"><h1>headline2.1</h1><div class=\"level-1\"><p>text2.1</p></div><h2>headline2.2</h2><div class=\"level-2\"><p>text2.2</p></div></div></div>",
		},
	}

	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}
}
