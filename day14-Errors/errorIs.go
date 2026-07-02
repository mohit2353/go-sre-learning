package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("record not found")

func database() error {
	return ErrNotFound
}

func service() error {
	err := database()
	if err != nil {
		return fmt.Errorf("service failed: %w", err)
	}
	return nil
}

func main() {
	err := service()

	if errors.Is(err, ErrNotFound) {
		fmt.Println("The Original error was record not found")
	}
}
