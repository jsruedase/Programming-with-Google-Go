package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Printf("Name: ")
	fmt.Scan(&name)
	fmt.Printf("Address: ")
	fmt.Scan(&address)

	mp := map[string]string{"name": name, "address": address}
	barr, _ := json.Marshal(mp)
	fmt.Printf(string(barr))

}
