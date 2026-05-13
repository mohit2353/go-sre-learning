package main

import "fmt"

func main() {

	var n int

	fmt.Print("Enter number: ")
	fmt.Scan(&n)

	fmt.Print("Even numbers: ")

	first := true

	for i := 1; i <= n; i++ {

		if i%2 == 0 {

			if !first {
				fmt.Print(", ")
			}

			fmt.Print(i)
			first = false
		}
	}

	fmt.Println()

	fmt.Print("Odd numbers: ")

	first = true

	for i := 1; i <= n; i++ {

		if i%2 != 0 {

			if !first {
				fmt.Print(", ")
			}

			fmt.Print(i)
			first = false
		}
	}

	fmt.Println()
}
