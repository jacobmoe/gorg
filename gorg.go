package gorg

import (
	"bufio"
	"fmt"
	"os"
)

func add(x int, y int) int {
	result := x + y
	return result
}

func scanFile(path string) {
	file, _ := os.Open(path)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
