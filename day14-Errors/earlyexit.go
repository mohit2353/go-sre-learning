package main

import (
	"errors"
	"fmt"
)

func authenticate() error {
	return errors.New("authencation failed")
}

func login() error {
	fmt.Println("Authenticating User....")

	err := authenticate()

	if err != nil {
		return err
	}

	fmt.Println("Welcome User!")

	return nil
}

func main() {

	fmt.Println("Starting the Login Process...")
	err := login()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Login Process Completed Successfully!")
}
