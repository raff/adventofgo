package main

import (
	"fmt"
	"sort"

	"github.com/raff/adventofgo/advlib"
)

type OpFunc func(old uint64) uint64

func AddFunc(v int) OpFunc {
	fmt.Println("ADD", v)
	vv := uint64(v)

	return func(old uint64) uint64 {
		if vv == 0 {
			return old + old
		}

		return old + vv
	}
}

func MulFunc(v int) OpFunc {
	fmt.Println("MUL", v)
	vv := uint64(v)

	return func(old uint64) uint64 {
		if vv == 0 {
			return old * old
		}

		return old * vv
	}
}

/*
func SubFunc(v int) OpFunc {
	fmt.Println("SUB", v)
        vv := uint64(v)

	return func(old uint64) uint64 {
		if vv == 0 {
			return 0
		}

		return old - vv
	}
}

func DivFunc(v int) OpFunc {
	fmt.Println("DIV", v)
        vv := uint64(v)

	return func(old uint64) uint64 {
		if vv == 0 {
			return 1
		}

		return old / vv
	}
}
*/

type Monkey struct {
	Name      int      // monkey number
	Initial   []int    // initial list of items
	Items     []uint64 // current list of items
	Op        OpFunc
	TestDiv   int // test if divisible by
	DestTrue  int // destination monkey
	DestFalse int // destination monkey

	Inspected int
}

func (m *Monkey) Reset() {
	m.Items = make([]uint64, len(m.Initial))

	for i, v := range m.Initial {
		m.Items[i] = uint64(v)
	}

	m.Inspected = 0
}

type ByInspected []Monkey

func (s ByInspected) Len() int           { return len(s) }
func (s ByInspected) Less(i, j int) bool { return s[i].Inspected < s[j].Inspected }
func (s ByInspected) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type ByName []Monkey

func (s ByName) Len() int           { return len(s) }
func (s ByName) Less(i, j int) bool { return s[i].Name < s[j].Name }
func (s ByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	r := advlib.NewReader()
	lines, _ := r.Readlines()

	var monkeys []Monkey

	var monkey Monkey

	for _, line := range lines {
		if line == "" {
			fmt.Println("EOM", monkey)
			monkeys = append(monkeys, monkey)
			continue
		}

		parts := advlib.Split(line)

		switch parts[0] {
		case "Monkey":
			n := advlib.ParseIntTrim(parts[1], ":")
			fmt.Println("MONKEY", n)
			monkey.Name = n

		case "Starting":
			items := advlib.ToInts(parts[2:])
			fmt.Println("ITEMS", items)
			monkey.Initial = items

		case "Operation:":
			op := parts[1:]
			fmt.Println("OPERATION", op)

			opcode := op[3]
			opval := advlib.ParseInt(op[4])

			switch opcode {
			case "+":
				monkey.Op = AddFunc(opval)

			case "-":
				//monkey.Op = SubFunc(opval)
				panic("no subtractions")

			case "*":
				monkey.Op = MulFunc(opval)

			case "/":
				//monkey.Op = DivFunc(opval)
				panic("no divisions")
			}

		case "Test:":
			test := parts[1:]
			fmt.Println("TEST", test)
			div := advlib.ParseInt(test[2])
			monkey.TestDiv = div

		case "If":
			cond := parts[1]
			stmt := parts[2:]
			fmt.Printf("IF %v THEN %q\n", cond, stmt)

			switch cond {
			case "true:":
				monkey.DestTrue = advlib.ParseInt(stmt[3])

			case "false:":
				monkey.DestFalse = advlib.ParseInt(stmt[3])
			}

		default:
			fmt.Printf("UNEXPECTED: %q\n", parts)
		}
	}

	fmt.Println("EOM", monkey)
	monkeys = append(monkeys, monkey)

	for part := 1; part <= 2; part++ {
		fmt.Println()
		fmt.Println("PART", part)

		sort.Sort(ByName(monkeys))

		for i, _ := range monkeys {
			m := &monkeys[i]
			m.Reset()
		}

		count := 20
		if part == 2 {
			count = 10000
		}

		for r := 0; r < count; r++ {
			fmt.Println("round", r+1)

			for i, monkey := range monkeys {
				//fmt.Println(monkey)

				for _, item := range monkey.Items {
					level := monkey.Op(item)
					if part == 1 {
						level /= 3
					}
					d := monkey.DestTrue
					if int(level)%monkey.TestDiv != 0 {
						d = monkey.DestFalse
					}
					monkeys[d].Items = append(monkeys[d].Items, level)

					//fmt.Println("level", level)
					//fmt.Println("div", monkey.TestDiv, (int(level) % monkey.TestDiv == 0))
					//fmt.Println("throw to", d)

					monkeys[i].Inspected++
				}

				monkeys[i].Items = monkeys[i].Items[:0]
				fmt.Println("monkey", monkey.Name, "/", i, "inspected", monkeys[i].Inspected)
			}
		}

		fmt.Println()
		fmt.Println("results")

		sort.Sort(sort.Reverse(ByInspected(monkeys)))

		for _, monkey := range monkeys {
			fmt.Printf("%#v\n", monkey)
		}

		fmt.Println()
		fmt.Println("monkey business:", monkeys[0].Inspected*monkeys[1].Inspected)
	}
}
