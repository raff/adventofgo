package main

import (
	"fmt"

	. "github.com/raff/adventofgo/advlib"
)

var (
	max = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func main() {
	r := NewReader()
	lines, _ := r.Readlines()

	total := 0

	//next_line:								; part 1
	for _, line := range lines {
		parts := SplitSep(line, ": ")

		_, game := Split2(parts[0])
		subs := SplitSep(parts[1], "; ")

		min := map[string]int{}

		for _, s := range subs {
			for _, c := range SplitSep(s, ", ") {
				n, color := Split2(c)
				in := ParseInt(n)
				/*						; part 1
				if in > max[color] {
					fmt.Println("skip game", game)
					continue next_line
				}
				*/
				if min[color] < in {
					min[color] = in
				}
			}
		}

		//fmt.Println(game, min)					; part 1
		//total += ParseInt(game)

		p := 1
		for _, v := range min {
			p *= v
		}

		fmt.Println(game, p)
		total += p
	}

	fmt.Println()
	fmt.Println("total", total)
}
