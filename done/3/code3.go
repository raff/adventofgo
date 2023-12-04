package main

import (
	"fmt"

	"github.com/gobs/matrix"
	. "github.com/raff/advent2022/advlib"
)

func main() {
	r := NewReader()
	lines, _ := r.Readlines()

	cols := len(lines[0])
	//rows := len(lines)

	m, _ := matrix.FromSlice(cols, false, []byte(Join(lines, "")))

        total := 0

	n := 0
	part := byte(0)
        partLoc := 0

        gears := map[int][]int{}    // part coordinats: part values

        addnum := func() {
		if n != 0 {
                        if part != 0 {
                            total += n
			    fmt.Println(n, string(part))

                            if part == '*' {
                                gears[partLoc] = append(gears[partLoc], n)
                            }
                        }

			n, part, partLoc = 0, 0, 0
		}
        }

	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			c := m.Get(x, y)
			if c >= '0' && c <= '9' {
				n = n*10 + int(c - '0')

				adj := m.Adjacent(x, y, false)
				for _, v := range adj {
					if (v.Value < '0' || v.Value > '9') && v.Value != '.' {
						part = v.Value
                                                partLoc = v.X * 1000 + v.Y
                                                break
					}
				}
			} else {
                            addnum()
			}
		}

                addnum()
	}

        addnum()

        fmt.Println()
        fmt.Println("total:", total)

        total = 0

        for _, v := range gears {
            if len(v) == 2 {
                r := v[0] * v[1]
                total += r
            }
        }

        fmt.Println("total ratios:", total)
}
