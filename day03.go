package main

import (
	"fmt"
)

const treeMark = "#"

// Day3 shows the answer.
func Day3() {
	lines, err := getLines("inputs/day03.txt")
	fatal(err)

	treesCount := 0
	hPos := 0

	for vPos := 1; vPos < len(lines); vPos++ {
		currentLine := lines[vPos]
		if (len(currentLine)-(hPos+3))-1 < 0 {
			hPos = -1 * (len(currentLine) - (hPos + 3))
		} else {
			hPos += 3
		}

		if string(currentLine[hPos]) == treeMark {
			treesCount++
		}
	}

	fmt.Printf("answer: %d\n", treesCount)
	// 145
}