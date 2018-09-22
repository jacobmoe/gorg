package gorg

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
)

type itemType int

const (
	itemErr itemType = iota
	itemEOF
	itemHeading
	itemText
	itemCode
)

type lexItem struct {
	typ itemType
	val string
	lev int
}

type lex struct {
	input *bufio.Scanner
	item  *lexItem
}

func (l *lex) Lex(yylval *yySymType) int {
	t := l.nextItem()

	switch t.typ {
	case itemEOF:
		return 0
	case itemErr:
		log.Println("lexer error:", t.val)
		return 0
	default:
		yylval.item = t
		return i
	}
}

func (l *lex) nextItem() *lexItem {
	if !l.input.Scan() {
		return l.lastItem()
	}

	if l.item == nil {
		l.item = newLexItem(l.input.Text())
		return l.nextItem()
	}

	switch l.item.typ {
	case itemHeading:
		current := l.item
		l.item = newLexItem(l.input.Text())

		return current
	case itemText:
		next := newLexItem(l.input.Text())

		if next.typ == itemText {
			l.combineItem(next)
			return l.nextItem()
		}

		current := l.item
		l.item = next

		return current
	default:
		panic("invalid lex item type")
	}
}

func (l *lex) lastItem() *lexItem {
	err := l.input.Err()
	if err != nil {
		return &lexItem{
			typ: itemErr,
			val: err.Error(),
		}
	}

	if l.item == nil {
		return &lexItem{typ: itemEOF}
	}

	last := l.item
	l.item = nil

	return last
}

func (l *lex) combineItem(next *lexItem) {
	l.item.val = fmt.Sprintf(
		"%s\n%s",
		l.item.val,
		next.val,
	)
}

var regHeading = regexp.MustCompile(`^(\*+)\ (.*)`)

func newLexItem(s string) *lexItem {
	m := regHeading.FindStringSubmatch(s)
	if len(m) > 2 {
		return &lexItem{
			typ: itemHeading,
			val: m[2],
			lev: len(m[1]),
		}
	}

	// also need to check for code block type

	return &lexItem{
		typ: itemText,
		val: s,
	}
}

func (x *lex) Error(s string) {
	log.Printf("parse error: %s", s)
}
