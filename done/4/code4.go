package main

import (
	"fmt"

	. "github.com/raff/advent2022/advlib"
)

type card struct {
	id   int
	wins []string
	nums []string
}

func (c card) CheckWins(process func(c card)) {
	for _, n := range c.nums {
		for _, w := range c.wins {
			if n == w {
				process(c)
			}
		}
	}
}

func main() {
	r := NewReader()
	lines, _ := r.Readlines()

	var cards []card

	for _, line := range lines {
		parts := SplitSep(line, ": ")
		_, cnum := Split2(parts[0])

		parts = SplitSep(parts[1], " | ")
		wins := Split(parts[0])
		nums := Split(parts[1])

		cards = append(cards, card{ParseInt(cnum), wins, nums})
	}

	if !Part2 { // part 1
		total := 0

		for _, c := range cards {
			points := 0

			c.CheckWins(func(c card) {
				if points == 0 {
					points = 1
				} else {
					points += points
				}
			})

			fmt.Println(c.id, c.wins, c.nums, points)
			total += points
		}

		fmt.Println()
		fmt.Println("total:", total)
		return
	}

	// part 2

	i := 0

	for {
		if i > len(cards) {
			break
		}

		c := cards[i]
		n := 1

		c.CheckWins(func(c card) {
			// fmt.Println("win", c.id)
			n++
		})

		for j := 1; j < n; j++ {
			fmt.Println("add card", c.id+j)
			cards = append(cards, cards[c.id+j-1])
		}

		i++
	}

	fmt.Println("total cards:", len(cards))
}
