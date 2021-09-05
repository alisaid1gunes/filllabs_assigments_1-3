package main

import (
	"fmt"
	"sort"
	"strings"
)

func customSort(array []string) []string {
	var withA []string
	var withOutA []string

	for _, item := range array {
		if strings.Contains(item, "a") {
			withA = append(withA, item)
		} else {
			withOutA = append(withOutA, item)
		}
	}
	sort.Slice(withA, func(i, j int) bool {
		return strings.Count(withA[i], "a") > strings.Count(withA[j], "a")
	})
	sort.Slice(withOutA, func(i, j int) bool {
		return len(withOutA[i]) > len(withOutA[j])
	})

	return append(withA, withOutA...)
}
func main() {
	strs := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	fmt.Println(customSort(strs))

}
