package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (an Animal) Eat() {
	fmt.Println(an.food)
}

func (an Animal) Move() {
	fmt.Println(an.locomotion)
}

func (an Animal) Speak() {
	fmt.Println(an.noise)
}

func main() {

	for {
		var animal Animal
		var name, info string
		fmt.Printf(">")
		fmt.Scan(&name, &info)
		if name == "cow" {
			animal = Animal{"grass", "walk", "moo"}
		} else if name == "bird" {
			animal = Animal{"worms", "fly", "peep"}
		} else if name == "snake" {
			animal = Animal{"mice", "slither", "hsss"}
		} else {
			fmt.Println("Enter a valid animal name.")
		}
		if info == "eat" {
			animal.Eat()
		} else if info == "move" {
			animal.Move()
		} else if info == "speak" {
			animal.Speak()
		}
	}
}
