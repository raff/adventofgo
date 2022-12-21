package advlib

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

func ParseIntTrim(s, t string) int {
        v, _ := strconv.Atoi(strings.TrimRight(s, t))
	return v
}

func ToInts(ls []string) (li []int) {
        for _, s := range ls {
            li = append(li, ParseIntTrim(s, ","))
        }

        return li
}

func Split(s string) []string {
	return strings.Fields(s)
}

func SplitSep(s, sep string) []string {
	return strings.Split(s, sep)
}

func StartsWith(s, start string) bool {
        return strings.HasPrefix(s, start)
}

func ContainsChar(s string, c rune) bool {
    return strings.Contains(s, string(c))
}
