package main

import (
	"fmt"

	"./advlib"
)

var (
	trees = [][]int{}
)

func check(x, y int) int {
	h := trees[y][x]

	// check right
	xx := x - 1
	for xx >= 0 && trees[y][xx] < h {
		xx--
	}

	if xx < 0 {
		return 1
	}

	// check left
	l := len(trees[y])
	xx = x + 1
	for xx < l && trees[y][xx] < h {
		xx++
	}

	if xx >= l {
		return 1
	}

	// check top
	yy := y - 1
	for yy >= 0 && trees[yy][x] < h {
		yy--
	}

	if yy < 0 {
		return 1
	}

	// check left
	l = len(trees)
	yy = y + 1
	for yy < l && trees[yy][x] < h {
		yy++
	}

	if yy >= l {
		return 1
	}

	return 0
}

func score(x, y int) int {
	vs := 1
	h := trees[y][x]

	// calculate right
	{
		s := 0
		for xx := x - 1; xx >= 0; xx-- {
			s++

			if trees[y][xx] >= h {
				break
			}
		}

		vs *= s
	}

	// calculate left
	{
		l := len(trees[y])
		s := 0
		for xx := x + 1; xx < l; xx++ {
			s++

			if trees[y][xx] >= h {
				break
			}
		}

		vs *= s
	}

	// calculate top
	{
		s := 0
		for yy := y - 1; yy >= 0; yy-- {
			s++

			if trees[yy][x] >= h {
				break
			}
		}

		vs *= s
	}

	// calculate left
	{
		l := len(trees)
		s := 0
		for yy := y + 1; yy < l; yy++ {
			s++

			if trees[yy][x] >= h {
				break
			}
		}

		vs *= s
	}

	return vs
}

func main() {
	r := advlib.NewReader()

	for {
		line, err := r.Readline()
		if err != nil {
			fmt.Println(err)
			break
		}

		row := make([]int, len(line))

		for i, c := range line {
			row[i] = int(c - '0')
		}

		trees = append(trees, row)
	}

	for _, r := range trees {
		fmt.Println(r)
	}

	fmt.Println()

	// part 1
	h := len(trees)
	w := len(trees[0])

	// start with edges (perimeter)
	total := w*2 + (h-2)*2

	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			total += check(x, y)
		}
	}

	fmt.Println("part 1:", total)
	fmt.Println()

	// part 2
	hs := 0

	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			s := score(x, y)
			fmt.Println("x:", x, "y:", y, "h:", s)

			if s > hs {
				hs = s
			}
		}
	}

	fmt.Println("part 2:", hs)
}
