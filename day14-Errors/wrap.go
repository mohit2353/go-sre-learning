package main

import (
	"errors"
	"fmt"
)

func database() error {
	return errors.New("connection refused")
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

	fmt.Println(err)
}
