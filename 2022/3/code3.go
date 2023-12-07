package main

import (
	"fmt"

	"github.com/raff/adventofgo/advlib"
)

func find(c1, c2 string) (rune, int) {
	for _, c := range c1 {
		for _, k := range c2 {
			if c == k {
				var p int

				if c >= 'a' && c <= 'z' {
					p = int(c - 'a' + 1)
				} else {
					p = int(c - 'A' + 27)
				}

				return c, p
			}
		}
	}

	panic("found nothing")
}

func find3(g []string) (rune, int) {
	for _, c0 := range g[0] {
		for _, c1 := range g[1] {
			if c0 == c1 {
				for _, c2 := range g[2] {
					if c0 == c2 {
						var p int

						if c0 >= 'a' && c0 <= 'z' {
							p = int(c0 - 'a' + 1)
						} else {
							p = int(c0 - 'A' + 27)
						}

						return c0, p
					}
				}
			}
		}
	}

	panic("found nothing")
}

func main() {
	r := advlib.NewReader()
	lines, _ := r.Readlines()

	// part 1
	total := 0

	for _, line := range lines {
		l := len(line)
		c1, c2 := line[:l/2], line[l/2:]
		c, p := find(c1, c2)
		total += p

		fmt.Println(string(c), p, c1, c2)
	}

	fmt.Println("part 1:", total)

	// part 2
	var group []string
	total = 0

	for _, line := range lines {
		if len(group) < 3 {
			group = append(group, line)

			if len(group) < 3 {
				continue
			}
		}

		c, p := find3(group)
		total += p

		fmt.Println(string(c), p, group)
		group = group[:0]
	}

	fmt.Println("part 2:", total)
}
