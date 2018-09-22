package main

import (
	"github.com/jacobmoe/gorg"
	"os"
)

func main() {
	f, _ := os.Open(
		"/Users/jacobmoe/code/go/src/github.com/jacobmoe/gorg/client/ex.org",
	)
	defer f.Close()

	gorg.Parse(f)
}
