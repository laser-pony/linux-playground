package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const startPattern = "15:02:25"
	const endPattern = "15:02:31"

	var lines []string
	start, end := -1, -1

	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.Contains(line, startPattern) && start == -1 {
			start = lineNum
		}
		if strings.Contains(line, endPattern) {
			end = lineNum
		}
		lineNum++
	}

	if start != -1 && end != -1 && start <= end {
		for i := start; i <= end && i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	} else {
		fmt.Fprintln(os.Stderr, "Start or end pattern not found or in wrong order.")
	}
}
