package main

import (
	"fmt"
)

func findMostRepeated(array []string) string {

	var biggest int = 0
	var bigIndex int = 0

	freq := make(map[string]int)

	for _, arItem := range array {
		if freq[arItem] == 0 {
			freq[arItem] = 1
		} else {
			freq[arItem]++
		}
	}

	for index, frItem := range array {
		if freq[frItem] >= biggest {
			biggest = freq[frItem]
			bigIndex = index

		}
	}

	return array[bigIndex]
}

func main() {
	
	x := []string{"apple","pie","apple","red","red","red"}
	
	fmt.Println("Most Repeated Item: ", findMostRepeated(x))

}
