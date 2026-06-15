package main

import "fmt"

func changeName(name string) {

	name = "GitLab"
}

func changeAge(age int) {
	age = 30
}

func increaseAge(age int) {
	age++
}

func main() {

	name := "Mohit"
	age := 25

	changeAge(age)
	increaseAge(age)

	changeName(name)

	fmt.Println(name)
	fmt.Println(age)
}
