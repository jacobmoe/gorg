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

func OrgToHtmlFile(orgPath string, htmlPath string) {
	byteHtml := []byte(OrgToHtml(orgPath))

	err := ioutil.WriteFile(htmlPath, byteHtml, 0644)
	check(err)
}

func OrgToHtml(orgPath string) string {
	tree := Tree{nodes: nodesFromFile(orgPath)}

	return tree.toHtml()
}

func nodesFromFile(path string) []*Node {
	file, err := os.Open(path)
	check(err)

	defer func() {
		check(file.Close())
	}()

	scanner := bufio.NewScanner(file)
	var node *Node
	var nodes []*Node
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

			node.parent = node.findParent(nodes)
			nodes = append(nodes, node)
		} else {
			if len(nodes) == 0 {
				nodes = []*Node{&Node{section: []string{line}}}
			} else {
				lastNode := nodes[len(nodes)-1]
				lastNode.section = append(lastNode.section, line)
			}
		}
	}

	return nodes
}
