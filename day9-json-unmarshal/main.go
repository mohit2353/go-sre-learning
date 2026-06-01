package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	Name        string `json:"name"`
	CPU         int    `json:"cpu"`
	Memory      int    `json:"memory"`
	Status      string `json:"status"`
	Environment string `json:"Environment"`
}

func main() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var server Server
	err = json.Unmarshal(data, &server)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Server Name:", server.Name)
	fmt.Println("CPU Cores:", server.CPU)
	fmt.Println("Memory:", server.Memory)
	fmt.Println("Status:", server.Status)
	fmt.Println("Environment:", server.Environment)

}
