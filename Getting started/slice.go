package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 3)
	stop := false
	i := 0
	for !stop {
		fmt.Printf("Enter an integer to add to the slice: ")
		var nums string
		fmt.Scan(&nums)
		num, err := strconv.Atoi(nums)
		if err == nil {
			if i < 3 {
				s[i] = num
			} else {
				s = append(s, num)
			}
			i += 1
		} else if nums != "X" {
			fmt.Println("Please type an integer or 'X' to exit")
		} else {
			fmt.Println("Quited successfully!")
			stop = true
		}
	}
	sort.Slice(s, func(i2, j int) bool {
		return s[i2] < s[j]
	})
	for _, v := range s {
		fmt.Println(v)
	}
}
