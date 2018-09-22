package main

import (
	"misc/gorgp"
	"os"
)

func main() {
	f, _ := os.Open(
		"/Users/jacobmoe/code/go/src/misc/gorgp/client/ex.org",
	)
	defer f.Close()

	gorgp.Parse(f)
}
