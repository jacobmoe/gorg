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
	tree := NewTree(nodesFromFile(orgPath))

	return tree.toHtml()
}

func OrgToJsonFile(orgPath string, jsonPath string) {
	byteJson := []byte(OrgToJson(orgPath))

	err := ioutil.WriteFile(jsonPath, byteJson, 0644)
	check(err)
}

func OrgToJson(orgPath string) string {
	tree := NewTree(nodesFromFile(orgPath))

	return tree.toJson()
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

			node = &Node{Headline: headline, Position: position}

			node.parent = node.findParent(nodes)
			nodes = append(nodes, node)
		} else {
			if len(nodes) == 0 {
				nodes = []*Node{&Node{Position: 1, Section: []string{line}}}
			} else {
				lastNode := nodes[len(nodes)-1]
				lastNode.Section = append(lastNode.Section, line)
			}
		}
	}

	return nodes
}
