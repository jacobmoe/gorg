package gorg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// func TestScanFile(t *testing.T) {
// 	fmt.Println("testing scanFile")
// 	scanFile("test/test_file.org")
// }

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
								text:     "text1.2",
							},
						},
					},
					Subtree{
						levels: []Level{
							Level{
								headline: "headline2.1",
								position: 1,
								text:     "text2.1",
							},
							Level{
								headline: "headline2.2",
								position: 2,
								text:     "text2.2",
							},
						},
					},
				},
			},
			out: "<div class=\"tree\"><div class=\"subtree\"><h1>headline1.1</h1><h2>headline1.2</h2><p class=\"level-2\">text1.2</p></div><div class=\"subtree\"><h1>headline2.1</h1><p class=\"level-1\">text2.1</p><h2>headline2.2</h2><p class=\"level-2\">text2.2</p></div></div>",
		},
	}

	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}
}
