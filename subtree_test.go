package gorg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddLevel(t *testing.T) {
	var s Subtree
	assert.Equal(t, len(s.levels), 0)

	level := Level{headline: "headline1", position: 1}
	s = s.addLevel(level)
	assert.Equal(t, len(s.levels), 1)

	level = Level{
		headline: "headline2",
		position: 2,
		text:     []string{"text2"},
	}
	s = s.addLevel(level)
	assert.Equal(t, len(s.levels), 2)

	expected := Subtree{
		levels: []Level{
			Level{headline: "headline1", position: 1},
			Level{
				headline: "headline2",
				position: 2,
				text:     []string{"text2"},
			},
		},
	}

	assert.Equal(t, s, expected)

}

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
						text:     []string{"the text2"},
					},
				},
			},
			out: "<div class=\"subtree\"><h1>the headline1</h1><h2>the headline2</h2><div class=\"level-2\"><p>the text2</p></div></div>",
		},
		{
			in: Subtree{
				levels: []Level{
					Level{
						headline: "headline1",
						position: 1,
						text:     []string{"text1"},
					},
					Level{
						headline: "headline2",
						position: 2,
						text:     []string{"text2"},
					},
				},
			},
			out: "<div class=\"subtree\"><h1>headline1</h1><div class=\"level-1\"><p>text1</p></div><h2>headline2</h2><div class=\"level-2\"><p>text2</p></div></div>",
		},
	}

	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}
}
