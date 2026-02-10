package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left, middle, right []int

	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num == pivot {
			middle = append(middle, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(quickSort(left), middle...), quickSort(right)...)
}

func generateRandomNumbers(n, max int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max) + 1
	}

	return arr
}

func main() {

	var size, maxValue int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&size)

	fmt.Print("Masukkan nilai maksimal (range 1-max): ")
	fmt.Scan(&maxValue)

	original := generateRandomNumbers(size, maxValue)

	fmt.Printf("\nArray Original (%d elemen):\n%v\n\n", size, original)

	sorted := quickSort(original)
	fmt.Println("Hasil Quick Sort:")
	fmt.Printf("%v\n", sorted)
}
