package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	mc := make(map[string]Cow)
	mb := make(map[string]Bird)
	ms := make(map[string]Snake)

	for {
		var command, name, specie string
		fmt.Printf(">")
		fmt.Scan(&command, &name, &specie)

		if command == "newanimal" {
			if specie == "cow" {
				mc[name] = Cow{"grass", "walk", "moo"}
				fmt.Printf("Created %s as a cow!\n", name)
			} else if specie == "bird" {
				mb[name] = Bird{"worms", "fly", "peep"}
				fmt.Printf("Created %s as a bird!\n", name)
			} else if specie == "snake" {
				ms[name] = Snake{"mice", "slither", "hsss"}
				fmt.Printf("Created %s as a snake!\n", name)
			} else {
				fmt.Println("Invalid species")
			}
		} else if command == "query" {
			action := specie
			if animal, ok := mc[name]; ok && action != "" {
				switch action {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("Invalid action")
				}
			} else if animal, ok := mb[name]; ok && action != "" {
				switch action {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("Invalid action")
				}
			} else if animal, ok := ms[name]; ok && action != "" {
				switch action {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("Invalid action")
				}
			} else {
				fmt.Println("Animal not found")
			}
		} else {
			fmt.Println("Invalid command")
		}
	}
}
