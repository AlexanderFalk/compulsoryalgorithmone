package main

import (
	"fmt"
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
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
			fmt.Println("Duplicated value: ", item)
			os.Exit(1)
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
}

func hypothesis(length int) bool {

}

func Generate_Sequence(n int) {
	var listOfNumbers []int
	var index = 0
	if len(listOfNumbers) == 0 {
		fmt.Println("Error")
	}
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		value := rand.Intn(n)
		fmt.Println(value)

		fmt.Println(index)
		listOfNumbers = append(listOfNumbers, value)
		index++
		dup_count(listOfNumbers)

	}

}

func main() {
	Generate_Sequence(100)
}
