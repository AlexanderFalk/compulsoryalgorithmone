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
			deck[i] = j
		}
	}
	fmt.Println("DONE!")
}

func main() {
	setup()
	fmt.Println(deck[])
}
