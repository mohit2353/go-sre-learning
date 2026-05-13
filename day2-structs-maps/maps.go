package main

import "fmt"

func main() {

	user := make(map[string]string)

	var name string
	var role string

	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	fmt.Print("Enter role: ")
	fmt.Scan(&role)

	user["name"] = name
	user["role"] = role

	fmt.Println(user)
}
