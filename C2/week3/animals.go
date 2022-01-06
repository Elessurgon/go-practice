package main

import "fmt"

type Animal struct {
	eat, locomotion, speak string
}

func (ani *Animal) Eat() string {
	return ani.eat
}
func (ani *Animal) Move() string {
	return ani.locomotion
}
func (ani *Animal) Speak() string {
	return ani.speak
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	fmt.Println("To Quit: Ctrl + C")
	for true {
		fmt.Print(">")
		var animal, method string
		var ani Animal
		fmt.Scanln(&animal, &method)
		fmt.Println(animal)
		fmt.Println(method)
		switch animal {
		case "cow":
			ani = cow
		case "bird":
			ani = bird
		case "snake":
			ani = snake
		default:
			fmt.Printf("Don't know animal %s\n", animal)
		}

		switch method {
		case "eat":
			fmt.Printf("%s eats %s\n", animal, ani.Eat())
		case "move":
			fmt.Printf("%s %ss\n", animal, ani.Move())
		case "speak":
			fmt.Printf("%s %ss\n", animal, ani.Speak())
		default:
			fmt.Printf("Don't know attribute %s for animal %s\n", method, animal)
		}
	}
}
