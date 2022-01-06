package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	m := make(map[string]string)

	fmt.Println("Enter a name")
	nm, _ := reader.ReadString('\n')
	fmt.Println("Enter an address")
	ad, _ := reader.ReadString('\n')

	m["name"] = strings.TrimRight(nm, "\r\n")
	m["address"] = strings.TrimRight(ad, "\r\n")

	json, _ := json.Marshal(m)
	fmt.Println(string(json))
}
