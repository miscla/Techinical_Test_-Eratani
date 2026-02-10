package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	return s == reversed
}

func main() {

	fmt.Print("Masukkan string: ")
	var input string
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println(input + " itu PALINDROME")
	} else {
		fmt.Println(input + " itu BUKAN palindrome")
	}
}
