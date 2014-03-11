package gorg

import (
	"path/filepath"
	"testing"
)

func TestCreateTree(t *testing.T) {
	path, _ := filepath.Abs("test/test_file.org")
	createTree(path)
}
