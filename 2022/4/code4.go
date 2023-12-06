package main

import (
	"fmt"

	"./advlib"
)

func fullOverlap(line string) bool {
	var m1, m2, n1, n2 int

	fmt.Sscanf(line, "%d-%d,%d-%d", &m1, &m2, &n1, &n2)
	//fmt.Println(m1, m2, n1, n2)

	if (m1 <= n1 && m2 >= n2) || (n1 <= m1 && n2 >= m2) {
		return true
	}

	return false
}

func partialOverlap(line string) bool {
	var m1, m2, n1, n2 int

	fmt.Sscanf(line, "%d-%d,%d-%d", &m1, &m2, &n1, &n2)
	//fmt.Println(m1, m2, n1, n2)

	if (m1 < n1 && m2 < n1) || (n1 < m1 && n2 < m1) {
		return false
	}

	return true
}

func main() {
	r := advlib.NewReader()
	lines, _ := r.Readlines()

	part1total := 0
	part2total := 0

	for _, line := range lines {
		if fullOverlap(line) {
			fmt.Println("full overlap", line)
			part1total++
		}
		if partialOverlap(line) {
			fmt.Println("partial overlap", line)
			part2total++
		}
	}

	fmt.Println("part 1 total", part1total, "overlaps")
	fmt.Println("part 2 total", part2total, "overlaps")
}
