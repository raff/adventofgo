package main

import (
	"fmt"
	"slices"

	. "github.com/raff/adventofgo/advlib"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOAK
	FullHouse
	FourOAK
	FiveOAK
)

func (ht HandType) String() string {
	switch ht {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOAK:
		return "Three of a Kind"
	case FullHouse:
		return "Full House"
	case FourOAK:
		return "Four of a Kind"
	case FiveOAK:
		return "Five of a Kind"
	}

	return "Bad hand type"
}

var cardvalue = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

func findType(hand string) HandType {
	var counts [13]int

	for _, c := range hand {
		v := cardvalue[c]
		counts[v]++
	}

	ht := HighCard

	for _, count := range counts {
		switch count {
		case 5:
			ht = FiveOAK
			break

		case 4:
			ht = FourOAK
			break

		case 3:
			if ht == OnePair {
				ht = FullHouse
			} else {
				ht = ThreeOAK
			}
			break

		case 2:
			switch ht {
			case ThreeOAK:
				ht = FullHouse

			case OnePair:
				ht = TwoPair

			default:
				ht = OnePair
			}
			break
		}
	}

	if Part2 {
		jc := counts[cardvalue['J']]

		switch jc {
		case 1:
			switch ht {
			case HighCard:
				return OnePair

			case OnePair:
				return ThreeOAK

			case TwoPair:
				return FullHouse

			case ThreeOAK:
				return FourOAK

			case FourOAK, FullHouse:
				return FiveOAK
			}

		case 2:
			switch ht {
			case OnePair:
				return ThreeOAK

			case TwoPair:
				return FourOAK

			case FullHouse:
				return FiveOAK
			}

		case 3:
			switch ht {
			case ThreeOAK:
				return FourOAK

			case FullHouse:
				return FiveOAK
			}

		case 4:
			return FiveOAK
		}
	}

	return ht
}

type Hand struct {
	cards string
	bid   int
	rank  HandType
}

func main() {
	r := NewReader()
	lines, _ := r.Readlines()

	if Part2 {
		cardvalue = map[rune]int{
			'A': 12,
			'K': 11,
			'Q': 10,
			'T': 9,
			'9': 8,
			'8': 7,
			'7': 6,
			'6': 5,
			'5': 4,
			'4': 3,
			'3': 2,
			'2': 1,
			'J': 0,
		}
	}

	var hands []Hand

	for _, line := range lines {
		parts := Split(line)
		if len(parts) != 2 {
			fmt.Println("bad format", line)
			return
		}

		hand := Hand{cards: parts[0], bid: ParseInt(parts[1]), rank: findType(parts[0])}
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.rank != b.rank {
			return int(a.rank - b.rank)
		}

		for i := 0; i < len(a.cards); i++ {
			ca, cb := rune(a.cards[i]), rune(b.cards[i])

			if ca != cb {
				return cardvalue[ca] - cardvalue[cb]
			}
		}

		return 0
	})

	total := 0

	for i, h := range hands {
		win := h.bid * (i + 1)

		fmt.Println(i+1, h.cards, h.rank, win)
		total += win
	}

	fmt.Println("total:", total)
}
