//go:generate goyacc -o parser.go parser.y

package gorg

import (
	"bufio"
	"io"
)

func Parse(r io.Reader) string {
	scanner := bufio.NewScanner(r)

	yyParse(&lex{input: scanner})

	return ""
}
