package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Printf("Word to be checked: ")
	fmt.Scan(&word)
	if (strings.HasPrefix(word, "i") || strings.HasPrefix(word, "I")) &&
		(strings.HasSuffix(word, "n") || strings.HasSuffix(word, "N")) &&
		(strings.Contains(word, "a") || strings.Contains(word, "A")) {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
