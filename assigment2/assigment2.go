package main

import (
	"fmt"
)

func recursion(n int) {
	if n == 13 {
		fmt.Println(9)
		return
	} else {
		fmt.Println(n - 7)
		recursion(n + 2)
	}
}

func main() {

	recursion(9)

}
