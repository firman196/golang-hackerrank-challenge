package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Printf("The input word is: %s\n", input)
	result := CamelCase(input)
	fmt.Printf("Camel Case : %d\n", result)
}

func CamelCase(input string) string {
	arrInput := strings.Split(input, "")
	formater := arrInput[:4]
	str := arrInput[4:]

	var result string
	for i := 0; i < len(str); i++ {
		if formater[0] == "S" {

		}
	}
}
