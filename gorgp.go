//go:generate goyacc -o parser.go parser.y

package gorgp

import (
	"bufio"
	"io"
)

func Parse(r io.Reader) string {
	scanner := bufio.NewScanner(r)

	yyParse(&lex{input: scanner})

	return ""
}
