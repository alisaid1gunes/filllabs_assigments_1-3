package main

import (
	"reflect"
	"testing"
)

func TestCustomSort(t *testing.T) {
	strs := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	exp := []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"}
    result := customSort(strs)
    expected := exp
    if reflect.DeepEqual(result, expected) != true {
        t.Errorf("add() test returned an unexpected result: got %v want %v", result, expected)
    }
}


