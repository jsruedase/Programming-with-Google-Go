package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	type Person struct {
		fname string
		lname string
	}

	var file_name string
	fmt.Printf("Enter the file name (with .txt): ")
	fmt.Scan(&file_name)

	file, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	file_content := string(file)
	lines := strings.Split(file_content, "\n")

	var sl []Person
	for _, line := range lines {
		n_s := strings.Fields(line)
		person := Person{n_s[0], n_s[1]}
		sl = append(sl, person)
	}

	for _, pers := range sl {
		fmt.Printf("Name: %s  LastName: %s\n", pers.fname, pers.lname)
	}
}
