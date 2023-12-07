package main

import (
	"fmt"

	"github.com/raff/adventofgo/advlib"
)

var (
	stacks = map[int][]string{}
	moves  = [][3]int{}
)

func pop(i int) string {
	l := len(stacks[i])
	if l <= 0 {
		panic(fmt.Sprintf("stack %v empty", i))
	}

	v := stacks[i][l-1]
	stacks[i] = stacks[i][:l-1]
	return v
}

func push(i int, v string) {
	stacks[i] = append(stacks[i], v)
}

func popn(i, n int) []string {
	l := len(stacks[i])
	if l-n < 0 {
		panic(fmt.Sprintf("stack %v underflow", i))
	}

	v := stacks[i][l-n:]
	stacks[i] = stacks[i][:l-n]
	return v
}

func pushn(i int, v []string) {
	stacks[i] = append(stacks[i], v...)
}

func main() {
	r := advlib.NewReader()

	for {
		line, err := r.Readline()
		if err != nil {
			fmt.Println(err)
			break
		}

		if line == "" || line[0] == '[' || line[0] == ' ' {
			continue
		}

		parts := advlib.Split(line)

		if parts[0] >= "0" && parts[0] <= "9" {
			stack := advlib.ParseInt(parts[0])
			stacks[stack] = parts[1:]
		}

		if parts[0] == "move" {
			stack := advlib.ParseInt(parts[1])
			from := advlib.ParseInt(parts[3])
			to := advlib.ParseInt(parts[5])
			moves = append(moves, [3]int{stack, from, to})
		}
	}

	fmt.Println("stacks")
	fmt.Println(stacks)

	fmt.Println()

	fmt.Println("moves")
	fmt.Println(moves)

	if false { // part 1
		for _, m := range moves {
			n, from, to := m[0], m[1], m[2]

			for i := 0; i < n; i++ {
				push(to, pop(from))
			}

			//fmt.Println("move", n, "from", from, "to", to)
			//fmt.Println(stacks)
		}

		var result string

		for i := 1; i <= len(stacks); i++ {
			result += pop(i)
		}

		fmt.Println("final")
		fmt.Println(stacks)

		fmt.Println("part 1:", result)

	} else { // part 2
		for _, m := range moves {
			n, from, to := m[0], m[1], m[2]

			pushn(to, popn(from, n))

			//fmt.Println("move", n, "from", from, "to", to)
			//fmt.Println(stacks)
		}

		var result string

		for i := 1; i <= len(stacks); i++ {
			result += pop(i)
		}

		fmt.Println("final")
		fmt.Println(stacks)

		fmt.Println("part 2:", result)
	}
}
