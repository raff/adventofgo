package main

import (
	"fmt"
	"strings"

	"github.com/raff/adventofgo/advlib"
)

func getnum(l string, first bool) (n int) {

	for len(l) > 0 {
		switch {
		case l[0] == '1':
			n = 1
			l = l[1:]

		case strings.HasPrefix(l, "one"):
			n = 1
			l = l[2:] // (e)ight

		case l[0] == '2':
			n = 2
			l = l[1:]

		case strings.HasPrefix(l, "two"):
			n = 2
			l = l[2:] // (o)ne

		case l[0] == '3':
			n = 3
			l = l[1:]

		case strings.HasPrefix(l, "three"):
			n = 3
			l = l[2:] // (e)ight

		case l[0] == '4':
			n = 4
			l = l[1:]

		case strings.HasPrefix(l, "four"):
			n = 4
			l = l[4:]

		case l[0] == '5':
			n = 5
			l = l[1:]

		case strings.HasPrefix(l, "five"):
			n = 5
			l = l[3:] // (e)ight

		case l[0] == '6':
			n = 6
			l = l[1:]

		case strings.HasPrefix(l, "six"):
			n = 6
			l = l[3:]

		case l[0] == '7':
			n = 7
			l = l[1:]

		case strings.HasPrefix(l, "seven"):
			n = 7
			l = l[4:] // (n)ine

		case l[0] == '8':
			n = 8
			l = l[1:]

		case strings.HasPrefix(l, "eight"):
			n = 8
			l = l[4:] // (t)wo or (t)hree

		case l[0] == '9':
			n = 9
			l = l[1:]

		case strings.HasPrefix(l, "nine"):
			n = 9
			l = l[3:] // (e)ight

		default:
			l = l[1:]
			continue
		}

		//fmt.Printf("n: %v line: %q\n", n, l)

		if first {
			return n
		}
	}

	return n
}

func main() {
	r := advlib.NewReader()
	lines, _ := r.Readlines()

	total := 0

	for _, line := range lines {
		first := getnum(line, true)
		last := getnum(line, false)

		if first < 0 || last < 0 {
			fmt.Println("invalid input:", line)
			return
		}

		n := first*10 + last
		fmt.Println(line, "-", n)
		total += n
	}

	fmt.Println("total:", total)
}
