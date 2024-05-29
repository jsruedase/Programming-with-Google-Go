package main

import (
	"fmt"
	"strconv"
)

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}

func main() {
	var sl []int
	for i := 0; i < 10; i++ {
		var nums string
		fmt.Printf("Enter an integer to add to the slice or 'X' to stop: ")
		fmt.Scan(&nums)
		if nums == "X" {
			break
		}
		num, _ := strconv.Atoi(nums)
		sl = append(sl, num)
	}
	BubbleSort(sl)

	for _, v := range sl {
		fmt.Print(v, " ")
	}

}
