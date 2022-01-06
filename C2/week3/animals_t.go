package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	for {
		animals := make(map[string]Animal)
		animals["cow"] = Animal{"grass", "walk", "moo"}
		animals["bird"] = Animal{"worms", "fly", "peep"}
		animals["snake"] = Animal{"mice", "slither", "hsss"}

		var input string
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("> ")
		if scanner.Scan() {
			input = scanner.Text()
		}

		input = strings.ToLower(input)
		if input == "q" {
			break
		}

		args := strings.Split(input, " ")
		animal := animals[args[0]]

		if (Animal{}) == animal {
			fmt.Println("invalid input for animal, try again")
			continue
		}

		switch args[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("invalid input for action, try again")
		}
	}
}
