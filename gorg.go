/*
 * Converts an org-mode file to a JSON object
 * The JSON is a recursive structure - it's subtrees all the way down
 * Subtrees contain nodes and more subtrees. 
 * Properties of a node:
 *   sections: Paragraphs and code snippets under a headline. Supports markdown.
 *   position: The asterisk count on a headline
 *   headline: The headline
*/

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

// read nodes from the file
// needs to be simplified
// bug: if a table ends the file, it will not be included
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
	var isCodeBlock bool
	var isTable bool

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`\A(\*+)\ (.*)`) // should use \S
		submatch := r.FindStringSubmatch(line)

		if len(submatch) > 1 {
			isBlock = false
			isCodeBlock = false
			isTable = false

			headline = submatch[2]
			position = len(submatch[1])

			node = &Node{Headline: headline, Position: position}

			node.parent = node.findParent(nodes)
			nodes = append(nodes, node)
		} else {
			codeStartReg := regexp.MustCompile(`\A(\#\+BEGIN_SRC)(.*)`)
			codeEndReg := regexp.MustCompile(`\A(\#\+END_SRC)`)

			tableReg := regexp.MustCompile(`\A\s*\|.*`)

			if codeStartReg.MatchString(line) {
				isCodeBlock = true
			} else if codeEndReg.MatchString(line) {
				isCodeBlock = false
			}

			isTable = tableReg.MatchString(line) && !isCodeBlock
			isBlock = isCodeBlock || isTable

			section += line

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
