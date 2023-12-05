package main

import (
	"fmt"
	"sort"

	. "github.com/raff/advent2022/advlib"
)

type Mapping struct {
	from, to, count int
}

type Mappings []Mapping

type Map struct {
	name     string
	mappings Mappings
}

func (m Map) Convert(s int) int {
	for _, m := range m.mappings {
		if s < m.from {
			return s
		}

		if s < m.from+m.count {
			return s - m.from + m.to
		}
	}

	return s
}

func (m Mappings) Len() int           { return len(m) }
func (m Mappings) Less(i, j int) bool { return m[i].from < m[j].from }
func (m Mappings) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func main() {
	r := NewReader()
	lines, _ := r.Readlines()

	var maps []Map
	var seeds []int

	var cur Map

	for _, line := range lines {
		switch {
		case StartsWith(line, "seeds: "):
			seeds = ToInts(Split(line[6:]))

		case line == "":
			if cur.name != "" {
				sort.Sort(cur.mappings)
				maps = append(maps, cur)
				cur.name = ""
				cur.mappings = nil
			}

		case EndsWith(line, " map:"):
			cur.name = line[:len(line)-5]

		default:
			parts := Split(line)
			if len(parts) != 3 {
				panic("invalid mapping: " + line)
			}

			vals := ToInts(parts)
			cur.mappings = append(cur.mappings, Mapping{from: vals[1], to: vals[0], count: vals[2]})
		}
	}

	if cur.name != "" {
		sort.Sort(cur.mappings)
		maps = append(maps, cur)
	}

	if Part2 {
		fmt.Println("part 2")
	}

	fmt.Println("seeds:", seeds)
	fmt.Println()
	for _, m := range maps {
		fmt.Println(m.name)

		for _, mm := range m.mappings {
			fmt.Println(" ", mm)
		}
	}

	nextseed := func() (n int) {
		if len(seeds) == 0 {
			return -1
		}

		n, seeds = seeds[0], seeds[1:]
		return
	}

	if Part2 {
		first := true

		nextseed = func() (n int) {
			if len(seeds) < 2 {
				return -1
			}

			start, count := seeds[0], seeds[1]
			if first && !Verbose {
				fmt.Println("seed:", start)
				first = false
			}

			n = start + count - 1
			if count <= 1 {
				seeds = seeds[2:]
				first = true
			} else {
				seeds[1]--
			}
			return
		}
	}

	loc := 0x7fffffffffffffff

	for {
		s := nextseed()
		if s < 0 {
			break
		}

		if Verbose {
			fmt.Println()
			fmt.Println("seed", s)
		}

		for _, m := range maps {
			c := m.Convert(s)

			if Verbose {
				fmt.Println(m.name, c)
			}

			s = c
		}

		if s < loc {
			loc = s
		}
	}

	fmt.Println()
	fmt.Println("lowest:", loc)
}
