package main

import (
	"fmt"
	"strings"
)

func main() {

	// --- Directions
	// Given a string, return a new string with the reversed
	// order of characters
	// --- Examples
	// 	reverse('apple') = 'leppa'
	// 	reverse('hello') = 'olleh'
	// 	reverse('Greetings!') = '!sgniteerG'

	fmt.Println("Reverse `hello`: ", reverse("hello"))
}

func reverse(str string) string {
	var reversed string

	split := strings.Split(str, "")

	for _, v := range split {
		reversed = v + reversed
	}

	return reversed
}
