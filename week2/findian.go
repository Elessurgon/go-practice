package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string")
	line, _ := in.ReadString('\n')
	var s string = strings.ToLower(line)
	if s[0] == 'i' && s[len(s)-3] == 'n' && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
