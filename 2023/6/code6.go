package main

import (
	"fmt"

	. "github.com/raff/adventofgo/advlib"
)

func win(t, d int) int {
	n := 0

	for i := 1; i <= t; i++ {
		v := i * (t - i)
		if v > d {
			n++
		}
	}

	return n
}

func main() {
	r := NewReader()
	lines, _ := r.Readlines()

	if len(lines) != 2 {
		fmt.Println("invalid line count")
		return
	}

	if Part1 {
		times := ToInts(Split(lines[0][9:])) // Time:
		dists := ToInts(Split(lines[1][9:])) // Distance:

		if len(times) != len(dists) {
			fmt.Println("mismatched")
			return
		}

		fmt.Println("time:", times)
		fmt.Println("dist:", dists)
		fmt.Println()

		merr := 1

		for i, t := range times {
			d := dists[i]
			w := win(t, d)

			fmt.Println(t, d, w)
			merr *= w
		}

		fmt.Println()
		fmt.Println("margin of error:", merr)
	} else {
		// Part 2

		time := ParseInt(Replace(lines[0][9:], " ", ""))
		dist := ParseInt(Replace(lines[1][9:], " ", ""))

		fmt.Println("time:", time)
		fmt.Println("dist:", dist)
		fmt.Println()

		wins := win(time, dist)
		fmt.Println("wins:", wins)
	}
}
