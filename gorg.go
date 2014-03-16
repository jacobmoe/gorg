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
	var node *Node
	var headline string
	var position int

	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`\A(\**)\ (.*)`)
		submatch := r.FindStringSubmatch(line)

		if len(submatch) > 1 {
			headline = submatch[2]
			position = len(submatch[1])
			node = &Node{headline: headline, position: position}

			tree.addNode(node)
		} else {
			if tree.isEmpty() {
				tree.addNode(&Node{section: []string{line}})
			} else {
				lastLevel := tree.lastNode()
				lastLevel.section = append(lastLevel.section, line)
			}
		}
	}

	return tree
}
