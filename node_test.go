package gorg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindParent(t *testing.T) {
	fmt.Println("==== Node TestFindParent")

	var node Node
	var expected *Node
	var parent *Node

	testTree := Tree{
		nodes: []*Node{
			&Node{headline: "headline1", position: 1},
			&Node{headline: "headline2", position: 2},
			&Node{headline: "headline3", position: 3},
			&Node{headline: "headline4", position: 4},
			&Node{headline: "headline5", position: 3},
			&Node{headline: "headline6", position: 4},
		},
	}

	node = Node{position: 1}
	parent = node.findParent(testTree)
	assert.Equal(t, parent, expected)

	node = Node{position: 2}
	parent = node.findParent(testTree)
	expected = testTree.nodes[0]
	assert.Equal(t, parent, expected)

	node = Node{position: 3}
	parent = node.findParent(testTree)
	expected = testTree.nodes[1]
	assert.Equal(t, parent, expected)

	node = Node{position: 4}
	parent = node.findParent(testTree)
	expected = testTree.nodes[4]
	assert.Equal(t, parent, expected)

	node = Node{position: 5}
	parent = node.findParent(testTree)
	expected = testTree.nodes[5]
	assert.Equal(t, parent, expected)
}

func TestNodeToHtml(t *testing.T) {
	fmt.Println("==== Node TestNodeToHtml ")

	var tests = []struct {
		in  Node
		out string
	}{
		{
			in: Node{
				headline: "the headline",
				position: 1,
				section:  []string{"the text"},
			},
			out: "<h1>the headline</h1><div class=\"level-1\"><p>the text</p></div>",
		},
		{
			in: Node{
				headline: "the headline3",
				position: 2,
			},
			out: "<h2>the headline3</h2>",
		},
		{
			in: Node{
				headline: "the headline3",
			},
			out: "<h1>the headline3</h1>",
		},
	}

	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}
}
