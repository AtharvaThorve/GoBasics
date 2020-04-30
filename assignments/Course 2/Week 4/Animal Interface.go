package main

import (
	"fmt"
)

type animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {
	food       string
	locomotion string
	noise      string
}

func (c cow) Eat() {
	fmt.Println(c.food)
}

func (c cow) Move() {
	fmt.Println(c.locomotion)
}

func (c cow) Speak() {
	fmt.Println(c.noise)
}

type bird struct {
	food       string
	locomotion string
	noise      string
}

func (b bird) Eat() {
	fmt.Println(b.food)
}

func (b bird) Move() {
	fmt.Println(b.locomotion)
}

func (b bird) Speak() {
	fmt.Println(b.noise)
}

type snake struct {
	food       string
	locomotion string
	noise      string
}

func (s snake) Eat() {
	fmt.Println(s.food)
}

func (s snake) Move() {
	fmt.Println(s.locomotion)
}

func (s snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	var x, y, z string
	fmt.Println("Enter your input")
	animalMap := make(map[string]animal)
	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&x, &y, &z)
		if err != nil {
			fmt.Println("Invalid input", err)
			continue
		}
		if x == "newanimal" {
			switch z {
			case "cow":
				animalMap[y] = cow{"grass", "walk", "moo"}
				fmt.Println("Created animal")
			case "bird":
				animalMap[y] = bird{"worms", "fly", "peep"}
				fmt.Println("Created animal")
			case "snake":
				animalMap[y] = snake{"mice", "slither", "hsss"}
				fmt.Println("Created animal")
			default:
				fmt.Println("Incorrect input")
			}
		} else if x == "query" {
			switch z {
			case "eat":
				animalMap[y].Eat()
			case "move":
				animalMap[y].Move()
			case "speak":
				animalMap[y].Speak()
			default:
				fmt.Println("Incorrect input")
			}
		}
	}
}
