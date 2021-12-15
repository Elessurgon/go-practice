package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var txt string

	fmt.Println("Name of text file with its extension (example: names.txt)")
	fmt.Scanln(&txt)
	f, err := os.Open(txt)
	if err != nil {
		log.Fatal(err)
	}
	var names []Person = make([]Person, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		ns := strings.Split(str, " ")
		p := Person{fname: ns[0], lname: ns[1]}
		names = append(names, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, n := range names {
		fmt.Printf("First name: %s, Last name: %s\n", n.fname, n.lname)
	}
}
