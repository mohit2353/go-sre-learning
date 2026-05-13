package main

import "fmt"

func main() {

	var text string

	fmt.Print("Enter string: ")
	fmt.Scan(&text)

	reversed := ""

	for i := len(text) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}

	fmt.Println("Reversed:", reversed)
}
