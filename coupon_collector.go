package main

import (
	"fmt"
	"math/rand"
	"time"
)

var harmonicValues []float64

func collect_coupons(n int) {
	var isCollected = make([]bool, n)
	var count int = 0
	var distinct int = 0

	rand.Seed(time.Now().UTC().UnixNano())
	for distinct < n {
		value := rand.Intn(n)
		//fmt.Println(value)
		count++

		//Every time a value doesn't match anything in the array, then we add it.
		//When the distinct is bigger than n, then we stop the program and print the count
		if !isCollected[value] {
			distinct++
			isCollected[value] = true
		}
	}

	fmt.Println("Total collected coupons: ", count)
}

func getHarmonicValues(number int) float64 {
	if harmonicValues[number] != 0 {
		return harmonicValues[number]
	}

	index := number - 1
	//Find last computed value
	for harmonicValues[index] == 0 {
		index--
	}

	for index <= number {
		harmonicValues[index] = float64(1)/float64(index) + harmonicValues[index-1]
		index++
	}

	return harmonicValues[number]

}

func main() {
	collect_coupons(4)
}
