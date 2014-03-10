package gorg

import (
	"bufio"
	"fmt"
	"os"
)

func scanFile(path string) {
	file, _ := os.Open(path)
	tree := createTree(bufio.NewScanner(file))
}

func createTree(scanner Scanner) Tree {
	var tree Tree
	var subtree Subtree
	var level Level
	var headline string
	var position int

	var isNextLevel bool

	for scanner.Scan() {
		r, _ := regexp.Compile(`\A(\**\) (.*)`)
		submatch := r.FindStringSubmatch(scanner.Text())
		if len(submatch) > 1 {
			headline = submatch[2]
			position = len(submatch[1])
			level = Level{headline: headline, position: position}

			isNextLevel = subtree.lastLevel.position < position
			subtree = subtree.addLevel(level)

			if !subtree.isEmpty() || !isNextLevel {
				tree.addSubtree(subtree)
			}

			level := Level{headline: submatch[1], position: 1}
			subtree = subtree.addLevel()
		} else {
			if subtree.isEmpty() {
				subtree.addLevel(Level{text: line})
			} else {
				text := subtree.lastLevel.text
				subtree.lastLevel.text = text + "<br />" + line
			}
		}
	}

	return tree.addSubtree(subtree)
}
