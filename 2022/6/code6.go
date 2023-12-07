package main

import (
	"fmt"

	"github.com/raff/adventofgo/advlib"
)

func marker(s string, n int) int {
outer:
	for i := 0; i < len(s)-n; i++ {
		p := s[i : i+n]

		//fmt.Println(p)

		counters := map[rune]int{}

		for _, c := range p {
			if counters[c] == 1 {
				continue outer
			}

			counters[c]++
		}

		return i + n

	}

	return -1
}

func main() {
	r := advlib.NewReader()
	lines, _ := r.Readlines()

	// part 1
	fmt.Println("part 1")

	for _, line := range lines {
		fmt.Println(line, marker(line, 4))
	}

	// part 2
	fmt.Println("part 2")

	for _, line := range lines {
		fmt.Println(line, marker(line, 14))
	}
}
