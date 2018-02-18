package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

// duplicate list
func dup_count(list []int) {
	duplicate_frequency := make(map[int]int)
	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicate_frequency[item]
		if exist {
			fmt.Println("Duplicated value: ", item)
			count := len(duplicate_frequency)
			hypothesis(float64(count))
			os.Exit(1)
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map

		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}

}

func hypothesis(length float64) bool {
	result := (math.Pi * 50) / 2
	hypoConf := math.Sqrt(result)

	newLength := math.Floor(length)
	newHypo := math.Floor(hypoConf)

	if newLength == newHypo {
		fmt.Printf("Hypothesis confirmed: %d = %d", newLength, newHypo)
		return true
	}
	fmt.Printf("Hypothesis denied: %d = %d", newLength, newHypo)
	return false
}

func Generate_Sequence(n int) {
	var listOfNumbers []int
	var index = 0

	rand.Seed(time.Now().UTC().UnixNano())
	for {
		value := rand.Intn(n - 1)
		fmt.Println(value)

		//fmt.Println(index)
		listOfNumbers = append(listOfNumbers, value)
		index++
		dup_count(listOfNumbers)

	}

}

func main() {

	Generate_Sequence(50)

}
