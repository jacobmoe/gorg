package gorg

import (
	"bufio"
	"os"
	"regexp"
)

func createTree(path string) Tree {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var tree Tree
	var subtree Subtree
	var level Level
	var headline string
	var position int

	var isNextLevel bool

	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`\A(\**)\ (.*)`)
		submatch := r.FindStringSubmatch(line)

		if len(submatch) > 1 {
			headline = submatch[2]
			position = len(submatch[1])
			level = Level{headline: headline, position: position}

			isNextLevel = subtree.lastLevel().position < position
			subtree = subtree.addLevel(level)

			if !subtree.isEmpty() || !isNextLevel {
				tree.addSubtree(subtree)
			}

			level = Level{headline: submatch[1], position: 1}
			subtree = subtree.addLevel(level)
		} else {
			if subtree.isEmpty() {
				subtree.addLevel(Level{text: []string{line}})
			} else {
				lastLevel := subtree.lastLevel()
				lastLevel.text = append(lastLevel.text, line)
			}
		}
	}

	return tree.addSubtree(subtree)
}
