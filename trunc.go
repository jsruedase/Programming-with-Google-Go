package main

import "fmt"

func main() {
	var num float32
	fmt.Printf("Float number: ")
	fmt.Scan(&num)
	trunc := int32(num)
	fmt.Printf("Truncated number: %d", trunc)
}
