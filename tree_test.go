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

	testTree = Tree{nodes: []*Node{&Node{Headline: "test"}}}

	assert.False(t, testTree.isEmpty())
}

func TestAddNode(t *testing.T) {
	fmt.Println("==== Tree TestAddNode")

	var tree Tree

	assert.True(t, tree.isEmpty())

	node1 := &Node{Headline: "test node1", Position: 1}
	tree.addNode(node1)

	assert.False(t, tree.isEmpty())
	assert.Equal(t, tree.nodes[0].Headline, "test node1")

	node2 := &Node{Headline: "test node2", Position: 3}
	tree.addNode(node2)
	assert.Equal(t, tree.nodes[0].Headline, "test node1")
	assert.Equal(t, tree.nodes[1].Headline, "test node2")

	node3 := &Node{Headline: "test node3", Position: 2}
	tree.addNode(node3)
	assert.Equal(t, tree.nodes[0].Headline, "test node1")
	assert.Equal(t, tree.nodes[1].Headline, "test node2")
	assert.Equal(t, tree.nodes[2].Headline, "test node3")

	node4 := &Node{Headline: "test node4", Position: 4}
	tree.addNode(node4)
	assert.Equal(t, tree.nodes[0].Headline, "test node1")
	assert.Equal(t, tree.nodes[1].Headline, "test node2")
	assert.Equal(t, tree.nodes[2].Headline, "test node3")
	assert.Equal(t, tree.nodes[3].Headline, "test node4")

	node5 := &Node{Headline: "test node5", Position: 5}
	tree.addNode(node5)
	assert.Equal(t, tree.nodes[0].Headline, "test node1")
	assert.Equal(t, tree.nodes[1].Headline, "test node2")
	assert.Equal(t, tree.nodes[2].Headline, "test node3")
	assert.Equal(t, tree.nodes[3].Headline, "test node4")
	assert.Equal(t, tree.nodes[4].Headline, "test node5")

	var parent *Node
	assert.Equal(t, tree.nodes[0].parent, parent)
	assert.Equal(t, tree.nodes[1].parent, node1)
	assert.Equal(t, tree.nodes[2].parent, node1)
	assert.Equal(t, tree.nodes[3].parent, node3)
	assert.Equal(t, tree.nodes[4].parent, node4)
}

func TestAddSubtree(t *testing.T) {
	fmt.Println("==== Tree TestSubtree")

	var tree Tree

	subtree1 := Tree{nodes: []*Node{&Node{Headline: "test"}}}
	tree.addSubtree(&subtree1)

	subtree2 := Tree{nodes: []*Node{&Node{Headline: "test"}}}
	tree.addSubtree(&subtree2)

	subtree3 := Tree{nodes: []*Node{&Node{Headline: "test"}}}
	tree.addSubtree(&subtree3)

	subtree4 := Tree{nodes: []*Node{&Node{Headline: "test"}}}
	tree.addSubtree(&subtree4)

	assert.Equal(t, tree.subtrees[0], &subtree1)
	assert.Equal(t, tree.subtrees[1], &subtree2)
	assert.Equal(t, tree.subtrees[2], &subtree3)
	assert.Equal(t, tree.subtrees[3], &subtree4)

	assert.Equal(t, tree.subtrees[0].parent, &tree)
	assert.Equal(t, tree.subtrees[1].parent, &tree)
	assert.Equal(t, tree.subtrees[2].parent, &tree)
	assert.Equal(t, tree.subtrees[3].parent, &tree)
}

func TestTreeToHtml(t *testing.T) {
	fmt.Println("==== Tree TestTreeToHtml")

	node1 := &Node{Headline: "headline1", Position: 1}
	node2 := &Node{
		Headline: "headline2",
		Position: 2,
		Section:  []string{"the section for node2"},
		parent:   node1,
	}
	node3 := &Node{Headline: "headline3", Position: 3, parent: node2}
	node4 := &Node{
		Headline: "headline4",
		Position: 4,
		Section:  []string{"the section for node4"},
		parent:   node3,
	}
	node5 := &Node{Headline: "headline5", Position: 3, parent: node2}
	node6 := &Node{
		Headline: "headline6",
		Position: 4,
		Section:  []string{"the section for node6", "some more text"},
		parent:   node5,
	}

	nodes := []*Node{node1, node2, node3, node4, node5, node6}
	tree := NewTree(nodes)
	html := tree.toHtml()

	var tests = []struct {
		in  string
		out string
	}{
		{
			in:  html,
			out: "<div class=\"subtree\"><h1>headline1</h1><h2>headline2</h2><div class=\"level-2\"><p>the section for node2</p></div><div class=\"subtree\"><h3>headline3</h3><h4>headline4</h4><div class=\"level-4\"><p>the section for node4</p></div></div><div class=\"subtree\"><h3>headline5</h3><h4>headline6</h4><div class=\"level-4\"><p>the section for node6</p><p>some more text</p></div></div></div>",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.in)
	}

}

func TestIndexOfNode(t *testing.T) {
	fmt.Println("==== Tree TestIndexOfNode")

	node1 := &Node{Headline: "headline1", Position: 1}
	node2 := &Node{Headline: "headline2", Position: 2}
	node3 := &Node{Headline: "headline3", Position: 3}
	node4 := &Node{Headline: "headline4", Position: 4}

	tree := Tree{nodes: []*Node{node1, node2, node3, node4}}

	assert.Equal(t, tree.indexOfNode(node4), 3)
	assert.Equal(t, tree.indexOfNode(node1), 0)
	assert.Equal(t, tree.indexOfNode(node2), 1)
	assert.Equal(t, tree.indexOfNode(node3), 2)
}

func TestDeleteNode(t *testing.T) {
	fmt.Println("==== Tree TestDeleteNode")

	node1 := &Node{Headline: "headline1", Position: 1}
	node2 := &Node{Headline: "headline2", Position: 2}
	node3 := &Node{Headline: "headline3", Position: 3}
	node4 := &Node{Headline: "headline4", Position: 4}

	tree := Tree{nodes: []*Node{node1, node2, node3, node4}}

	assert.Equal(t, len(tree.nodes), 4)

	tree.deleteNode(node2)

	assert.Equal(t, len(tree.nodes), 3)
	assert.Equal(t, tree.nodes[0].Headline, "headline1")
	assert.Equal(t, tree.nodes[1].Headline, "headline3")
	assert.Equal(t, tree.nodes[2].Headline, "headline4")

	tree.deleteNode(node1)

	assert.Equal(t, len(tree.nodes), 2)
	assert.Equal(t, tree.nodes[0].Headline, "headline3")
	assert.Equal(t, tree.nodes[1].Headline, "headline4")

	tree.deleteNode(node4)

	assert.Equal(t, len(tree.nodes), 1)
	assert.Equal(t, tree.nodes[0].Headline, "headline3")

	tree.deleteNode(node3)

	assert.Equal(t, len(tree.nodes), 0)
}

func TestUnflattenTree(t *testing.T) {
	fmt.Println("==== Tree TestUnflattenTree")

	var tree Tree

	tree.addNode(&Node{Headline: "sub0.0.1", Position: 2})
	tree.addNode(&Node{Headline: "sub1.1.1", Position: 1})
	tree.addNode(&Node{Headline: "sub1.1.2", Position: 2})
	tree.addNode(&Node{Headline: "sub1.1.3", Position: 3})
	tree.addNode(&Node{Headline: "sub2.1.1", Position: 1})
	tree.addNode(&Node{Headline: "sub2.1.2", Position: 3})
	tree.addNode(&Node{Headline: "sub2.2.1", Position: 4})
	tree.addNode(&Node{Headline: "sub2.2.2", Position: 5})
	tree.addNode(&Node{Headline: "sub2.3.1", Position: 4})
	tree.addNode(&Node{Headline: "sub2.3.2", Position: 5})
	tree.addNode(&Node{Headline: "sub2.3.3", Position: 6})
	tree.addNode(&Node{Headline: "sub2.4.1", Position: 4})
	tree.addNode(&Node{Headline: "sub2.4.2", Position: 5})
	tree.addNode(&Node{Headline: "sub3.1.1", Position: 1})

	//  ** sub0.1.1
	//  * sub1.1.1
	//  ** sub1.1.2
	//  *** sub1.1.3
	//  * sub2.1.1
	//  *** sub2.1.2
	//  **** sub2.2.1
	//  ***** sub2.2.2
	//  **** sub2.3.1
	//  ***** sub2.3.2
	//  ****** sub2.3.3
	//  **** sub2.4.1
	//  ***** sub2.4.2
	//  * sub3.1.1

	tree.unflatten()

	assert.Equal(t, len(tree.nodes), 0)
	assert.Equal(t, len(tree.subtrees), 4)

	sub01 := tree.subtrees[0]
	assert.Equal(t, len(sub01.nodes), 1)
	assert.Equal(t, len(sub01.subtrees), 0)

	assert.Equal(t, sub01.nodes[0].Headline, "sub0.0.1")

	sub11 := tree.subtrees[1]
	assert.Equal(t, len(sub11.nodes), 3)
	assert.Equal(t, len(sub11.subtrees), 0)

	assert.Equal(t, sub11.nodes[0].Headline, "sub1.1.1")
	assert.Equal(t, sub11.nodes[1].Headline, "sub1.1.2")
	assert.Equal(t, sub11.nodes[2].Headline, "sub1.1.3")

	sub21 := tree.subtrees[2]
	assert.Equal(t, len(sub21.nodes), 2)
	assert.Equal(t, len(sub21.subtrees), 3)

	assert.Equal(t, sub21.nodes[0].Headline, "sub2.1.1")
	assert.Equal(t, sub21.nodes[1].Headline, "sub2.1.2")

	sub22 := tree.subtrees[2].subtrees[0]
	assert.Equal(t, len(sub22.nodes), 2)
	assert.Equal(t, len(sub22.subtrees), 0)

	assert.Equal(t, sub22.nodes[0].Headline, "sub2.2.1")
	assert.Equal(t, sub22.nodes[1].Headline, "sub2.2.2")

	sub23 := tree.subtrees[2].subtrees[1]
	assert.Equal(t, len(sub23.nodes), 3)
	assert.Equal(t, len(sub23.subtrees), 0)

	assert.Equal(t, sub23.nodes[0].Headline, "sub2.3.1")
	assert.Equal(t, sub23.nodes[1].Headline, "sub2.3.2")
	assert.Equal(t, sub23.nodes[2].Headline, "sub2.3.3")

	sub24 := tree.subtrees[2].subtrees[2]
	assert.Equal(t, len(sub24.nodes), 2)
	assert.Equal(t, len(sub24.subtrees), 0)

	assert.Equal(t, sub24.nodes[0].Headline, "sub2.4.1")
	assert.Equal(t, sub24.nodes[1].Headline, "sub2.4.2")

	sub31 := tree.subtrees[3]
	assert.Equal(t, len(sub31.nodes), 1)
	assert.Equal(t, len(sub31.subtrees), 0)

	assert.Equal(t, sub31.nodes[0].Headline, "sub3.1.1")
}
