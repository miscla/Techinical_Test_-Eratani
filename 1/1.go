package main

import (
	"fmt"
)

func isDisliked(n int) bool {
	if n%3 == 0 {
		return true
	}

	if n%10 == 3 {
		return true
	}

	return false
}

func findLiked(k int) int {
	count := 0
	num := 0

	for count < k {
		// println(num)
		// println(!isDisliked(num))
		num++
		if !isDisliked(num) {
			// println(num)
			// println("masuk")
			count++
		}
	}

	return num
}

func main() {
	var t int
	fmt.Scan(&t)

	results := make([]int, t)

	for i := 0; i < t; i++ {
		var k int
		fmt.Scan(&k)
		results[i] = findLiked(k)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
