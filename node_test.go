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

	nodes := []*Node{
		&Node{Headline: "headline1", Position: 1},
		&Node{Headline: "headline2", Position: 2},
		&Node{Headline: "headline3", Position: 3},
		&Node{Headline: "headline4", Position: 4},
		&Node{Headline: "headline5", Position: 3},
		&Node{Headline: "headline6", Position: 4},
	}

	node = Node{Position: 1}
	parent = node.findParent(nodes)
	assert.Equal(t, parent, expected)

	node = Node{Position: 2}
	parent = node.findParent(nodes)
	expected = nodes[0]
	assert.Equal(t, parent, expected)

	node = Node{Position: 3}
	parent = node.findParent(nodes)
	expected = nodes[1]
	assert.Equal(t, parent, expected)

	node = Node{Position: 4}
	parent = node.findParent(nodes)
	expected = nodes[4]
	assert.Equal(t, parent, expected)

	node = Node{Position: 5}
	parent = node.findParent(nodes)
	expected = nodes[5]
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
				Headline: "the headline",
				Position: 1,
				Section:  []string{"the text"},
			},
			out: "<h1>the headline</h1><div class=\"level-1\"><p>the text</p></div>",
		},
		{
			in: Node{
				Headline: "the headline3",
				Position: 2,
			},
			out: "<h2>the headline3</h2>",
		},
		{
			in: Node{
				Headline: "the headline3",
				Position: 4,
			},
			out: "<h4>the headline3</h4>",
		},
	}

	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}
}

func TestNodeToJson(t *testing.T) {
	fmt.Println("==== Node TestJsonToHtml")

	node := Node{
		Headline: "the headline",
		Position: 2,
		Section:  []string{"the text", "more text", "even more"},
	}

	expected := "{\"headline\":\"the headline\",\"position\":2,\"sections\":[\"the text\",\"more text\",\"even more\"]}"

	assert.Equal(t, expected, node.toJson())
}
