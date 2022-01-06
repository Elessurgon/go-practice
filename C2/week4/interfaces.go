package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	s := make(map[string]Animal)
	for true {
		fmt.Print("Please enter a command:\n\tnewanimal <name> <type>\n\tquery <name> <action>\n")
		var input string
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print(">")
		if scanner.Scan() {
			input = scanner.Text()
		}
		args := strings.Split(input, " ")

		if len(args) == 3 && args[0] == "newanimal" {
			if args[2] == "cow" || args[2] == "snake" || args[2] == "bird" {
				switch args[2] {
				case "cow":
					s[args[1]] = Cow{}
				case "bird":
					s[args[1]] = Bird{}
				case "snake":
					s[args[1]] = Snake{}
				}
				fmt.Println("Created it")
			} else {
				fmt.Println("Not created")
			}
		}
		if len(args) == 3 && args[0] == "query" {
			_, ok := s[args[1]]
			if ok {
				switch args[2] {
				case "eat":
					s[args[1]].Eat()
				case "move":
					s[args[1]].Move()
				case "speak":
					s[args[1]].Speak()
				default:
					fmt.Println("Undefined behaviour")
				}
			} else {
				fmt.Println("Name doesn't Exist")
			}
		}
	}
}
