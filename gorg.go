package gorg

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func OrgToHtml(orgPath string, htmlPath string) {
	var html string

	html = createTree(orgPath).toHtml()

	byteHtml := []byte(html)

	err := ioutil.WriteFile(htmlPath, byteHtml, 0644)
	check(err)
}

func createTree(path string) Tree {
	file, err := os.Open(path)
	check(err)

	defer func() {
		check(file.Close())
	}()

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
