package main

import (
	"fmt"
)

func GenDisplaceFn(a, v_0, s_0 float64) func(float64) float64 {

	fn := func(t float64) float64 {
		return 0.5*a*t*t + v_0*t + s_0
	}
	return fn
}

func main() {
	var a, v_0, s_0, t float64
	fmt.Printf("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Printf("Enter initial velocity: ")
	fmt.Scan(&v_0)
	fmt.Printf("Enter initial displacement: ")
	fmt.Scan(&s_0)
	fmt.Printf("Enter time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v_0, s_0)
	fmt.Println(fn(t))
}
