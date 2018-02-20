package main

import (
	"fmt"
)

type Rank int

const (
	One Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Suit int

const (
	Spades Suit = iota + 1
	Hearts
	Clubs
	Diamonds
)

var deck = make(map[Suit]Rank)

func setup() {

	for i := Spades; i <= Diamonds; i++ {
		for j := One; j <= King; j++ {
			deck = append(deck, i)
			deck = append(deck, j)
		}
	}
	fmt.Println("DONE!")
}

func shellsort(items []int) {
	var (
		n    = len(items)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if items[j-gap] > items[j] {
					items[j-gap], items[j] = items[j], items[j-gap]
				}
				j = j - gap
			}
		}
	}
}

func main() {
	setup()
	fmt.Println()
}
