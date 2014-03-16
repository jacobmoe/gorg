package gorg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	fmt.Println("==== Tree TestisEmpty")

	var testTree Tree

	assert.True(t, testTree.isEmpty())

	testTree = Tree{nodes: []*Node{&Node{headline: "test"}}}

	assert.False(t, testTree.isEmpty())
}

func TestAddNode(t *testing.T) {
	fmt.Println("==== Tree TestAddNode")

	var tree Tree

	assert.True(t, tree.isEmpty())

	node1 := &Node{headline: "test node1", position: 1}
	tree.addNode(node1)

	assert.False(t, tree.isEmpty())
	assert.Equal(t, tree.nodes[0].headline, "test node1")

	node2 := &Node{headline: "test node2", position: 3}
	tree.addNode(node2)
	assert.Equal(t, tree.nodes[0].headline, "test node1")
	assert.Equal(t, tree.nodes[1].headline, "test node2")

	node3 := &Node{headline: "test node3", position: 2}
	tree.addNode(node3)
	assert.Equal(t, tree.nodes[0].headline, "test node1")
	assert.Equal(t, tree.nodes[1].headline, "test node2")
	assert.Equal(t, tree.nodes[2].headline, "test node3")

	node4 := &Node{headline: "test node4", position: 4}
	tree.addNode(node4)
	assert.Equal(t, tree.nodes[0].headline, "test node1")
	assert.Equal(t, tree.nodes[1].headline, "test node2")
	assert.Equal(t, tree.nodes[2].headline, "test node3")
	assert.Equal(t, tree.nodes[3].headline, "test node4")

	node5 := &Node{headline: "test node5", position: 5}
	tree.addNode(node5)
	assert.Equal(t, tree.nodes[0].headline, "test node1")
	assert.Equal(t, tree.nodes[1].headline, "test node2")
	assert.Equal(t, tree.nodes[2].headline, "test node3")
	assert.Equal(t, tree.nodes[3].headline, "test node4")
	assert.Equal(t, tree.nodes[4].headline, "test node5")

	var parent *Node
	assert.Equal(t, tree.nodes[0].parent, parent)
	assert.Equal(t, tree.nodes[1].parent, node1)
	assert.Equal(t, tree.nodes[2].parent, node1)
	assert.Equal(t, tree.nodes[3].parent, node3)
	assert.Equal(t, tree.nodes[4].parent, node4)
}

func TestTreeToHtml(t *testing.T) {
	node1 := &Node{headline: "headline1", position: 1}
	node2 := &Node{
		headline: "headline2",
		position: 2,
		parent:   node1,
		section:  []string{"the section for node2"},
	}
	node3 := &Node{headline: "headline3", position: 3, parent: node2}
	node4 := &Node{
		headline: "headline4",
		position: 4,
		parent:   node3,
		section:  []string{"the section for node4"},
	}
	node5 := &Node{headline: "headline5", position: 3, parent: node2}
	node6 := &Node{
		headline: "headline6",
		position: 4,
		parent:   node5,
		section:  []string{"the section for node6", "some more text"},
	}

	var tests = []struct {
		in  Tree
		out string
	}{
		{
			in: Tree{
				nodes: []*Node{node1, node2, node3, node4, node5, node6},
			},
			out: "<div class=\"tree\"><h1>headline1</h1><h2>headline2</h2><div class=\"level-2\"><p>the section for node2</p></div><h3>headline3</h3><h4>headline4</h4><div class=\"level-4\"><p>the section for node4</p></div><h3>headline5</h3><h4>headline6</h4><div class=\"level-4\"><p>the section for node6</p><p>some more text</p></div></div>",
		},
	}

	for _, test := range tests {
		actual := test.in.toHtml()
		assert.Equal(t, test.out, actual)
	}

}
