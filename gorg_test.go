package gorg

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestCreateTree(t *testing.T) {
	fmt.Println("==== Tree testCreateTree")

	path, _ := filepath.Abs("test/test_file.org")
	tree := createTree(path)
	fmt.Println(tree.toHtml())
}
