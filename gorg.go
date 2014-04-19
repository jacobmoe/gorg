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

func TreeFromFile(orgPath string) *Tree {
	return NewTree(nodesFromFile(orgPath))
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

func OrgToJson(orgPath string) []byte {
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
	var section string
	var isBlock bool

	for scanner.Scan() {
		line := scanner.Text()

		r, _ := regexp.Compile(`\A([^\ ]\**)\ (.*)`) // should use \S
		submatch := r.FindStringSubmatch(line)

		if len(submatch) > 1 {
			isBlock = false
			headline = submatch[2]
			position = len(submatch[1])

			node = &Node{Headline: headline, Position: position}

			node.parent = node.findParent(nodes)
			nodes = append(nodes, node)
		} else {
			codeBlockStartReg, _ := regexp.Compile(`\A(\#\+BEGIN_SRC)(.*)`)
			codeBlockEndReg, _ := regexp.Compile(`\A(\#\+END_SRC)`)

			if codeBlockStartReg.MatchString(line) {
				isBlock = true
			} else if codeBlockEndReg.MatchString(line) {
				isBlock = false
			}

			section += (line + "\n")

			if !isBlock {
				if len(nodes) == 0 {
					nodes = []*Node{&Node{Position: 1, Section: []string{section}}}
				} else {
					lastNode := nodes[len(nodes)-1]
					lastNode.Section = append(lastNode.Section, section)
				}
				section = ""
			}
		}
	}

	return nodes
}
