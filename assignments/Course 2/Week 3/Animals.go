package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	var x, y string
	fmt.Printf("Enter animal name and type of action\nEnter x and x for both inputs to quit\n")
	for {
		fmt.Print("> ")
		fmt.Scan(&x, &y)
		strings.ToLower(x)
		strings.ToLower(y)
		if x == "x" && y == "x" {
			fmt.Println("Goodbye :)")
			break
		} else {
			var a Animal
			switch x {
			case "cow":
				a = cow
			case "bird":
				a = bird
			case "snake":
				a = snake
			default:
				fmt.Print("Enter valid input for x ")
			}
			switch y {
			case "eat":
				a.Eat()
			case "speak":
				a.Speak()
			case "move":
				a.Move()
			default:
				fmt.Println("Enter valid input for y")
			}
		}
	}

}
