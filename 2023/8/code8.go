package main

import (
	"fmt"

	. "github.com/raff/adventofgo/advlib"
)

type node struct {
	name  string
	left  string
	right string
}

type instructions struct {
	seq string
	cur int
}

func (i *instructions) reset() {
	i.cur = 0
}

func (i *instructions) next() (ret byte) {
	ret = i.seq[i.cur]

	i.cur++

	if i.cur >= len(i.seq) {
		i.cur = 0
	}

	return
}

func main() {
	r := NewReader()
	lines, _ := r.Readlines()

	var instr *instructions
	nodes := map[string]node{}
	starts := []string{}

	for _, line := range lines {
		if instr == nil {
			instr = &instructions{seq: line}
			continue
		}
		if line == "" {
			continue
		}

		parts := Split(line)
		name := parts[0]
		// = -> parts[1]
		left := Trim(parts[2], " =(),")
		right := Trim(parts[3], " =(),")

		n := node{name, left, right}
		nodes[name] = n

		if EndsWith(name, "A") {
			starts = append(starts, name)
		}
	}

	if Verbose {
		fmt.Println()
		fmt.Println("instructions:", instr)
		fmt.Println("nodes:")
		for k, v := range nodes {
			fmt.Println(" ", k, v)
		}
	}

	fmt.Println()

	if Part1 {
		count := 0

		for cur := "AAA"; cur != "ZZZ"; {
			dir := instr.next()
			count++

			n := nodes[cur]
			if dir == 'R' {
				cur = n.right
			} else {
				cur = n.left
			}

			fmt.Println(count, cur, n, string(dir))
		}

		return
	}

	// part 2

	// there is an easier method here that is to calculate
	// the paths one by one (for start in range starts -> calculate path)
	// and then calculate the least common denominator of the results

	count := 0
	max := 0

	for {
		count++

		dir := instr.next()
		ends := 0

		for i, s := range starts {
			n := nodes[s]
			if dir == 'R' {
				s = n.right
			} else {
				s = n.left
			}

			if EndsWith(s, "Z") {
				ends++
			}

			starts[i] = s
		}

		if ends > max {
			max = ends
			fmt.Println(count, starts, max)
		} else if count%10000000 == 0 {
			fmt.Println(count, starts, max)
		}
		if ends == len(starts) {
			break
		}
	}
}
