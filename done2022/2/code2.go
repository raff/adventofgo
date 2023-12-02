package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func (r *Reader) Readlines() ([]string, error) {
	var lines []string

	for {
		line, err := r.Readline()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("ERROR", err)
			return nil, err
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func ParseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

const (
	Lose = 0
	Draw = 3
	Win  = 6
)

var (
	// part 1
	shapes = map[string]int{
		"X": 1, // rock
		"Y": 2, // paper
		"Z": 3, // scissor
	}

	scores = map[string]int{
		"XA": Draw, // rock, rock
		"XB": Lose, // rock, paper
		"XC": Win,  // rock, scissor
		"YA": Win,  // paper, rock
		"YB": Draw, // paper, paper
		"YC": Lose, // paper, scissor
		"ZA": Lose, // scissor, rock
		"ZB": Win,  // scissor, paper
		"ZC": Draw, // scissor, scissor
	}

	// part 2
	results = map[string]int{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}

	choices = map[string]string{
		"AX": "Z", // lose, scissor
		"AY": "X", // draw, rock
		"AZ": "Y", // win, paper
		"BX": "X", // lose, rock
		"BY": "Y", // draw, paper
		"BZ": "Z", // win, scissor
		"CX": "Y", // lose, paper
		"CY": "Z", // draw, scissor
		"CZ": "X", // win, rock
	}
)

func main() {
	r := NewReader()

	lines, err := r.Readlines()
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	// part 1

	total := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		fmt.Println(parts)

		sc := parts[1] + parts[0]
		total += shapes[parts[1]] + scores[sc]
	}

	fmt.Println("A:", total)

	// part 2
	total = 0

	for _, line := range lines {
		parts := strings.Fields(line)

		sc := parts[0] + parts[1]
		c := choices[sc]

		fmt.Println(parts, shapes[c], choices[sc], results[parts[1]])

		total += shapes[c] + results[parts[1]]
	}

	fmt.Println("B:", total)
}
