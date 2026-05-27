package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name   string `json:"name"`
	CPU    int    `json:"cpu"`
	Memory int    `json:"memory"`
	Status string `json:"status"`
}

func main() {

	server1 := Server{
		Name:   "Production-1",
		CPU:    75,
		Memory: 90,
		Status: "Working",
	}

	jsonData, err := json.Marshal(server1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
