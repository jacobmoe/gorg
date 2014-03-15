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
