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

func (animal Cow) Eat() {
	fmt.Println("grass")
}

func (animal Cow) Move() {
	fmt.Println("walk")
}

func (animal Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (animal Bird) Eat() {
	fmt.Println("worms")
}

func (animal Bird) Move() {
	fmt.Println("fly")
}

func (animal Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (animal Snake) Eat() {
	fmt.Println("mice")
}

func (animal Snake) Move() {
	fmt.Println("slither")
}

func (animal Snake) Speak() {
	fmt.Println("hiss")
}

func NewAnimalCommand(animals *map[string]Animal, name, animal string) {
	switch animal {
	case "cow":
		(*animals)[name] = Cow{}
	case "bird":
		(*animals)[name] = Bird{}
	case "snake":
		(*animals)[name] = Snake{}
	default:
		fmt.Println("No such animal type")
	}
}

func QueryCommand(animals *map[string]Animal, name, action string) {
	animal, ok := (*animals)[name]
	if !ok {
		fmt.Println("No animal found by that name")
		return
	}

	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Action must be one of (eat, move, speak)")
	}
}

func main() {
	animals := make(map[string]Animal)

	for {
		fmt.Print("\n> ")

		reader := bufio.NewReader(os.Stdin)
		buffer, _, _ := reader.ReadLine()
		tokens := strings.Split(string(buffer), " ")
		if len(tokens) != 3 {
			fmt.Println("Requests must have 3 parts")
			continue
		}

		action := strings.ToLower(tokens[0])
		arg1 := strings.ToLower(tokens[1])
		arg2 := strings.ToLower(tokens[2])

		switch action {
		case "newanimal":
			NewAnimalCommand(&animals, arg1, arg2)
		case "query":
			QueryCommand(&animals, arg1, arg2)
		default:
			fmt.Println("Query must be either \"newanimal\" or \"query\"")
		}
	}
}
