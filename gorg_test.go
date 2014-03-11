package gorg

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestCreateTree(t *testing.T) {
	path, _ := filepath.Abs("test/test_file.org")
	tree := createTree(path)
	fmt.Println(tree.toHtml())
}
