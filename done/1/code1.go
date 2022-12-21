package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Reader struct {
	sc  *bufio.Scanner
	err error
}

func NewReader() *Reader {
	return &Reader{sc: bufio.NewScanner(os.Stdin)}
}

func (r *Reader) Readline() (string, error) {
	if r.err != nil {
		return "", r.err
	}

	if r.sc.Scan() {
		return r.sc.Text(), nil
	}

	r.err = r.sc.Err()
	if r.err == nil {
		r.err = io.EOF
	}

	r.sc = nil
	return "", r.err
}

func ParseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func main() {
	r := NewReader()

	var cals []int
	var curr int

	for {
		line, err := r.Readline()
		if err == io.EOF {
			if curr > 0 {
				cals = append(cals, curr)
			}

			break
		}

		if err != nil {
			fmt.Println("ERROR", err)
			return
		}

		if line == "" {
			cals = append(cals, curr)
			curr = 0
		} else {
			curr += ParseInt(line)
		}
	}

	fmt.Println("Data:", cals)
	sort.Sort(sort.Reverse(sort.IntSlice(cals)))
	fmt.Println("Sorted:", cals)
	fmt.Println("max", cals[0], cals[1], cals[2], "=>", cals[0]+cals[1]+cals[2])
}
