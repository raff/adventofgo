package main

import . "github.com/raff/adventofgo/advlib"

func predict(seq []int, prev bool) int {
	l := make([]int, len(seq))
	copy(l, seq)

	ll := [][]int{l}

	for {
		var d []int

		sum := 0

		for i := 1; i < len(l); i++ {
			dd := l[i] - l[i-1]
			sum += dd
			d = append(d, dd)
		}

		l = d
		ll = append(ll, l)

		if sum == 0 {
			break
		}
	}

	v := 0

	for i := len(ll) - 1; i >= 0; i-- {
		l := ll[i]

		if prev {
			p := l[0]
			v = p - v
			l = append([]int{v}, l...)
			ll[i] = l
		} else {
			p := l[len(l)-1]
			v += p
			l = append(l, v)
			ll[i] = l
		}

		Println(l)
	}

	return v
}

func main() {
	r := NewReader()
	lines, _ := r.Readlines()

	var seqs [][]int

	for _, line := range lines {
		seq := ToInts(Split(line))
		seqs = append(seqs, seq)
	}

	total := 0

	for _, s := range seqs {
		p := predict(s, Part2)
		Println(p)

		total += p
	}

	Println("total:", total)
}
