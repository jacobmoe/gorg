package gorg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddLevel(t *testing.T) {
	fmt.Println("==== TestAddLevel")

	var s Subtree
	assert.Equal(t, len(s.levels), 0)

	level1 := Level{headline: "headline1", position: 1}
	s.addLevel(&level1)
	assert.Equal(t, len(s.levels), 1)

	level2 := Level{
		headline: "headline2",
		position: 2,
		text:     []string{"text2"},
	}
	s.addLevel(&level2)
	assert.Equal(t, len(s.levels), 2)

	expected := Subtree{levels: []*Level{&level1, &level2}}

	assert.Equal(t, s, expected)
}

func TestAddBranch(t *testing.T) {
	fmt.Println("==== TestAddBranch")

	var s Subtree
	assert.Equal(t, len(s.branches), 0)

	subtree1 := Subtree{
		levels: []*Level{
			&Level{headline: "headline1", position: 1},
		},
	}
	s.addBranch(&subtree1)
	assert.Equal(t, len(s.branches), 1)

	subtree2 := Subtree{
		levels: []*Level{
			&Level{headline: "headline2.1", position: 1},
			&Level{headline: "headline2.2", position: 2},
		},
	}

	s.addBranch(&subtree2)
	assert.Equal(t, len(s.branches), 2)
	assert.Equal(t, subtree1.parent, &s)
	assert.Equal(t, subtree2.parent, &s)

	expected := Subtree{branches: []*Subtree{&subtree1, &subtree2}}

	assert.Equal(t, s, expected)
}

func TestRoot(t *testing.T) {
	fmt.Println("==== TestRoot")

	var s Subtree

	assert.Equal(t, *s.root(), s)

	sub := Subtree{
		levels: []*Level{
			&Level{},
		},
	}

	s.addBranch(&sub)

	subsub := Subtree{
		levels: []*Level{
			&Level{},
		},
	}

	sub.addBranch(&subsub)

	assert.Equal(t, &s, s.branches[0].parent)
	assert.Equal(t, &s, sub.root())

	assert.Equal(t, &sub, s.branches[0].branches[0].parent)
	assert.Equal(t, &s, subsub.root())
}

func TestLastLevel(t *testing.T) {
	fmt.Println("==== TestLastLevel")

	var subtree Subtree

	assert.Nil(t, subtree.lastLevel())

	level := Level{
		headline: "headline1",
		position: 1,
		text:     []string{"text1"},
	}

	subtree.addLevel(&level)
	assert.Equal(t, *subtree.lastLevel(), level)

	level = Level{
		headline: "headline2",
		position: 2,
		text:     []string{"text2"},
	}

	subtree.addLevel(&level)
	assert.Equal(t, *subtree.lastLevel(), level)
}

func TestIsEmpty(t *testing.T) {
	fmt.Println("==== TestIsEmpty")

	var subtree Subtree
	assert.True(t, subtree.isEmpty())

	level := Level{
		headline: "headline1",
		position: 1,
		text:     []string{"text1"},
	}

	subtree.addLevel(&level)
	assert.False(t, subtree.isEmpty())
}

func TestSubtreeToHtml(t *testing.T) {
	fmt.Println("==== TestSubtreeToHtml")

	var tests = []struct {
		in  Subtree
		out string
	}{
		{
			in: Subtree{
				levels: []*Level{
					&Level{
						headline: "the headline1",
						position: 1,
					},
					&Level{
						headline: "the headline2",
						position: 2,
						text:     []string{"the text2"},
					},
				},
			},
			out: "<div class=\"subtrees\"><h1>the headline1</h1><h2>the headline2</h2><div class=\"level-2\"><p>the text2</p></div></div>",
		},
		{
			in: Subtree{
				levels: []*Level{
					&Level{
						headline: "headline1",
						position: 1,
						text:     []string{"text1"},
					},
					&Level{
						headline: "headline2",
						position: 2,
						text:     []string{"text2"},
					},
				},
			},
			out: "<div class=\"subtrees\"><h1>headline1</h1><div class=\"level-1\"><p>text1</p></div><h2>headline2</h2><div class=\"level-2\"><p>text2</p></div></div>",
		},
	}

	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}
}
