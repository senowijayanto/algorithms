package main

import (
	"fmt"
	"strings"
)

func main() {
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
