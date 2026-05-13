package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	var user1 User

	fmt.Print("Name: ")
	fmt.Scan(&user1.Name)

	fmt.Print("Age: ")
	fmt.Scan(&user1.Age)

	fmt.Println("User", user1.Name, "is", user1.Age, "years old")
}
