package main

import (
	"fmt"

	"./advlib"
)

type Grid struct {
	content []byte
	width   int
	height  int
}

func NewGrid(w, h int) (g Grid) {
	g.width = w
	g.height = h
        g.content = make([]byte, w * h)
        return g
}

func (g Grid) Reset(c byte) {
        for i := range g.content {
                    g.content[i] = c
	}
}

func (g Grid) Print() {
	p := 0

	for y := 0; y < g.height; y++ {
		fmt.Println(string(g.content[p : p+g.width]))
		p += g.width
	}
}

func (g Grid) Set(x, y int, c byte) {
	// y = g.height + y - 1
	x += g.width / 2
	y += g.height / 2

	fmt.Println("set", x, y)

	g.content[y*g.width+x] = c
}

func (g Grid) Count(c byte) (count int) {
	for _, b := range g.content {
		if b == c {
			count++
		}
	}

	return count
}

type Pos struct {
	x, y int
}

type Tail struct {
	list []Pos
	max  int
}

func NewTail(l int) *Tail {
	return &Tail{max: l}
}

func (t *Tail) Add(p Pos) bool {
	t.list = append(t.list, p)
	l := len(t.list)

	if l >= t.max {
		t.list = t.list[l-t.max:]
                return true
	}

        return false
}

func (t *Tail) First() Pos {
	l := len(t.list)
	return t.list[l-1]
}

func (t *Tail) Last() Pos {
	return t.list[0]
}

func main() {
	r := advlib.NewReader()
	lines, _ := r.Readlines()

	//grid := NewGrid(500, 800)
        grid := NewGrid(40, 40)
        ltail := NewTail(9)

	for part := 2; part <= 2; part++ {
	        grid.Reset('.')

                h := Pos{0, 0}
                t := Pos{0, 0}

		for _, line := range lines {
			parts := advlib.Split(line)
			dir := parts[0]
			steps := advlib.ParseInt(parts[1])

			fmt.Println(dir, steps)

			for i := 0; i < steps; i++ {
				switch dir {
				case "R": // right
					h.x++
					if h.x-t.x == 2 {
						t.x++
						t.y = h.y
					}
				case "L": // left
					h.x--
					if t.x-h.x == 2 {
						t.x--
						t.y = h.y
					}
				case "D": // down
					h.y++
					if h.y-t.y == 2 {
						t.y++
						t.x = h.x
					}
				case "U": // up
					h.y--
					if t.y-h.y == 2 {
						t.y--
						t.x = h.x
					}
				}

				fmt.Printf("H:%d,%d, T:%d,%d\n", h.x, h.y, t.x, t.y)

                                if part == 1 {
				        grid.Set(t.x, t.y, '#')
                                } else if ltail.Add(t) {
                                        last := ltail.Last()
				        grid.Set(last.x, last.y, '#')

                                        fmt.Println("last", last)
                                }

			}

                        grid.Print()
		}

		fmt.Println("part", part, "total:", grid.Count('#'))
	}
}
