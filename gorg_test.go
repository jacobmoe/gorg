package gorg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestNodesFromFile(t *testing.T) {
	fmt.Println("==== gorg TestNodesFromFile")
	path, _ := filepath.Abs("test/test.org")

	tree := NewTree(nodesFromFile(path))

	assert.Equal(t, 0, len(tree.Nodes))
	assert.Equal(t, 4, len(tree.Subtrees))

	expected := "<div class=\"subtree\"><div class=\"level-1\"><p>body 0</p></div></div><div class=\"subtree\"><h1>headline 1.1</h1><h2>headline 1.2</h2><h3>headline 1.3</h3><h4>headline 1.4</h4><div class=\"level-4\"><p>body 1</p></div></div><div class=\"subtree\"><h1>headline 2.1</h1><div class=\"level-1\"><p>body 2.1.1</p><p>body 2.1.2</p></div><h2>headline 2.2</h2><div class=\"level-2\"><p>body 2.2.1</p><p>body 2.2.2</p></div><h3>headline 2.3</h3><div class=\"level-3\"><p>body 2</p><p>body 2.2</p></div></div><div class=\"subtree\"><h1>headline 3.1</h1><div class=\"level-1\"><p>  body 3.1</p><p>body 3.2</p><p>body 3.3</p></div></div>"

	assert.Equal(t, tree.toHtml(), expected)
}

func TestOrgToHtmlFile(t *testing.T) {
	fmt.Println("==== gorg testOrgToHtmlFile")

	inPath, _ := filepath.Abs("test/test.org")
	outPath, _ := filepath.Abs("test/test.html")

	// remove last test file, if exists
	if _, err := os.Stat(outPath); err == nil {
		os.Remove(outPath)
	}

	OrgToHtmlFile(inPath, outPath)

	htmlFileContents, _ := ioutil.ReadFile(outPath)
	contents := htmlFileContents

	expected := "<div class=\"subtree\"><div class=\"level-1\"><p>body 0</p></div></div><div class=\"subtree\"><h1>headline 1.1</h1><h2>headline 1.2</h2><h3>headline 1.3</h3><h4>headline 1.4</h4><div class=\"level-4\"><p>body 1</p></div></div><div class=\"subtree\"><h1>headline 2.1</h1><div class=\"level-1\"><p>body 2.1.1</p><p>body 2.1.2</p></div><h2>headline 2.2</h2><div class=\"level-2\"><p>body 2.2.1</p><p>body 2.2.2</p></div><h3>headline 2.3</h3><div class=\"level-3\"><p>body 2</p><p>body 2.2</p></div></div><div class=\"subtree\"><h1>headline 3.1</h1><div class=\"level-1\"><p>  body 3.1</p><p>body 3.2</p><p>body 3.3</p></div></div>"

	assert.Equal(t, string(contents), expected)
}

func TestOrgToJsonFile(t *testing.T) {
	fmt.Println("==== gorg testOrgToJsonFile")

	inPath, _ := filepath.Abs("test/test.org")
	outPath, _ := filepath.Abs("test/test.json")

	// remove last test file
	if _, err := os.Stat(outPath); err == nil {
		os.Remove(outPath)
	}

	OrgToJsonFile(inPath, outPath)

	jsonFileContents, _ := ioutil.ReadFile(outPath)
	contents := jsonFileContents

	expected := "{\"nodes\":[],\"subtrees\":[{\"nodes\":[{\"headline\":\"\",\"position\":1,\"sections\":[\"body 0\"]}],\"subtrees\":null},{\"nodes\":[{\"headline\":\"headline 1.1\",\"position\":1,\"sections\":null},{\"headline\":\"headline 1.2\",\"position\":2,\"sections\":null},{\"headline\":\"headline 1.3\",\"position\":3,\"sections\":null},{\"headline\":\"headline 1.4\",\"position\":4,\"sections\":[\"body 1\"]}],\"subtrees\":null},{\"nodes\":[{\"headline\":\"headline 2.1\",\"position\":1,\"sections\":[\"body 2.1.1\",\"body 2.1.2\"]},{\"headline\":\"headline 2.2\",\"position\":2,\"sections\":[\"body 2.2.1\",\"body 2.2.2\"]},{\"headline\":\"headline 2.3\",\"position\":3,\"sections\":[\"body 2\",\"body 2.2\"]}],\"subtrees\":null},{\"nodes\":[{\"headline\":\"headline 3.1\",\"position\":1,\"sections\":[\"  body 3.1\",\"body 3.2\",\"body 3.3\"]}],\"subtrees\":null}]}"

	assert.Equal(t, string(contents), expected)
}
