package main

import "fmt"

func main() {
	// --- Directions
	// Given a string, return true if the string is a palindrome
	// or false if it is not.  Palindromes are strings that
	// form the same word if it is reversed. *Do* include spaces
	// and punctuation in determining if the string is a palindrome.
	// --- Examples:
	//   palindrome("abba") = true
	//   palindrome("abcdefg") = false
	fmt.Println("Is abba palindrome: ", palindrome("abba"))
}

func palindrome(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
