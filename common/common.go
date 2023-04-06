package common

import (
	"bufio"
	"strings"
)

func ConsoleColumnWidth(text string, finalColWidth int) int {
	return finalColWidth - len(text)
}

// Return slice with each line of a multi-line string, splitting on '\n'
func SplitStringLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
