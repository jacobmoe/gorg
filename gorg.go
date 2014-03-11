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

	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`\A(\**)\ (.*)`)
		submatch := r.FindStringSubmatch(line)

		if len(submatch) > 1 {
			headline = submatch[2]
			position = len(submatch[1])
			level = Level{headline: headline, position: position}

			subtree.addLevel(&level)

			if subtree.lastLevel().position < position {
				tree.addSubtree(subtree)
			}
		} else {
			if subtree.isEmpty() {
				subtree.addLevel(&Level{text: []string{line}})
			} else {
				lastLevel := subtree.lastLevel()
				lastLevel.text = append(lastLevel.text, line)
			}
		}
	}

	return tree.addSubtree(subtree)
}
